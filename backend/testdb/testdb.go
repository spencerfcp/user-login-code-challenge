package testdb

import (
	"context"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	// has the side effect of loading the postgres driver
	_ "github.com/lib/pq"
)

func Open(fn func(*sqlx.DB)) {
	databaseURL := os.Getenv("TEST_DATABASE_URL")
	if err := open(fn, databaseURL, true); err != nil {
		panic(err)
	}
}

func open(fn func(*sqlx.DB), databaseUrl string, clearData bool) error {
	db, err := sqlx.Open("postgres", databaseUrl)
	if err != nil {
		return errors.Wrap(err, "failed to open database")
	}

	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	// prevent multiple tests from running at the same time
	conn, err := db.Conn(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get database connection")
	}

	defer func() {
		if err = conn.Close(); err != nil {
			panic(errors.Wrap(err, "failed to close testdb locking connection"))
		}
	}()

	_, err = conn.ExecContext(context.Background(), "SELECT pg_advisory_lock(1)")
	if err != nil {
		return errors.Wrap(err, "failed to get lock")
	}
	defer func() {
		_, err = conn.ExecContext(context.Background(), "SELECT pg_advisory_unlock(1)")
		if err != nil {
			panic(errors.Wrap(err, "failed to release lock"))
		}
	}()

	// One connection for our advisary lock and one connection for our tests
	db.SetMaxOpenConns(2)

	if clearData {
		_, err := db.Exec(`
			delete from users;
		`)

		if err != nil {
			return errors.Wrap(err, "failed to cleanup test data")
		}
	}

	fn(db)

	return nil
}

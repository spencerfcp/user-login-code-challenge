package db_user

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/spencerfcp/scoirapi/v2/database"
	"github.com/spencerfcp/scoirapi/v2/id"
	"github.com/spencerfcp/scoirapi/v2/ptr"
	"golang.org/x/crypto/bcrypt"
)

type UserRow struct {
	ID         id.UserID
	Password   string
	Username   id.Username
	Created_at time.Time
	Updated_at time.Time
}

type Row struct {
	Username id.Username
	Password string
}

func UsernameExists(db sqlx.Ext, username id.Username) bool {
	return database.Exists(db, `
		select 1 from users where lower(username) = lower($1)
	`, username)
}

func Add(db sqlx.Ext, user Row) (id.UserID, error) {
	uId, err := database.InsertStructReturningID(db, "users", "id", user)
	return id.UserID(uId), err
}

func GetUser(db sqlx.Ext, userID id.UserID) (UserRow, error) {
	return get(db, &userID, nil)
}

func GetUserByUsername(db sqlx.Ext, username id.Username) (*UserRow, error) {
	if username == "" {
		return nil, nil
	}

	result, err := get(db, nil, &username)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.Wrapf(err, "GetUserByUsername failed for %s", username)
	}

	return &result, nil
}

func get(
	db sqlx.Ext,
	optionalUserID *id.UserID,
	optionalUsername *id.Username,
) (UserRow, error) {
	userID := ptr.WithDefault(optionalUserID, -1)
	username := ptr.WithDefault(optionalUsername, "")

	var user UserRow
	err := sqlx.Get(db, &user, `
		select
			id
			, username
			, password
			, created_at
			, updated_at
		from
			users
		where
			($1 = -1 or $1 = id)
			and ($2 = '' or lower($2) = lower(username))
	`, userID, username)

	return user, err
}

func HashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

func PasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))

	return err == nil
}

package db_user_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/spencerfcp/user-login-code-challenge/backend/database/db_user"
	"github.com/spencerfcp/user-login-code-challenge/backend/testdb"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByUsername(t *testing.T) {
	testdb.Open(func(db *sqlx.DB) {
		testUser :=
			db_user.Row{
				Username: "random",
				Password: "WhatAGreatPassword!!1",
			}

		uID, err := db_user.Add(db, db_user.Row{
			Username: testUser.Username,
			Password: testUser.Password,
		})
		assert.NoError(t, err)

		user, err := db_user.GetUserByUsername(db, "Random")
		assert.NoError(t, err)

		assert.Equal(t, db_user.UserRow{
			ID:       uID,
			Username: testUser.Username,
			Password: testUser.Password,
		}, db_user.UserRow{
			ID:       user.ID,
			Password: user.Password,
			Username: user.Username,
		})
	})
}

func TestGetUser(t *testing.T) {
	testdb.Open(func(db *sqlx.DB) {
		testUser :=
			db_user.Row{
				Username: "random",
				Password: "WhatAGreatPassword!!1",
			}

		uID, err := db_user.Add(db, db_user.Row{
			Username: testUser.Username,
			Password: testUser.Password,
		})
		assert.NoError(t, err)

		user, err := db_user.GetUser(db, uID)
		assert.NoError(t, err)
		assert.Equal(t, db_user.UserRow{
			ID:       uID,
			Username: testUser.Username,
			Password: testUser.Password,
		}, db_user.UserRow{
			ID:       user.ID,
			Password: user.Password,
			Username: user.Username,
		})
	})
}

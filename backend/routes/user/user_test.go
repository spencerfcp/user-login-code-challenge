package user_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/spencerfcp/user-login-code-challenge/backend/database/db_user"
	"github.com/spencerfcp/user-login-code-challenge/backend/handler"
	"github.com/spencerfcp/user-login-code-challenge/backend/pb"
	"github.com/spencerfcp/user-login-code-challenge/backend/routes/user"
	"github.com/spencerfcp/user-login-code-challenge/backend/testdb"
	"github.com/spencerfcp/user-login-code-challenge/backend/testreq"
	"github.com/stretchr/testify/assert"
)

func TestUserLogin(t *testing.T) {
	testdb.Open(func(db *sqlx.DB) {
		testUser :=
			db_user.Row{
				Username: "random",
				Password: "WhatAGreatPassword!!1",
			}

		_, err := user.InsertUser(handler.UnauthenticatedParams{
			Message: &pb.UserSignupRequest{
				Username: string(testUser.Username),
				Password: string(testUser.Password),
			},
			DB: db,
		})
		assert.NoError(t, err)

		result, err := user.Login(handler.UnauthenticatedParams{
			Message: &pb.UserLoginRequest{
				Username: string(testUser.Username),
				Password: string(testUser.Password),
			},
			DB: db,
		})
		assert.NoError(t, err)

		loginUserResponse := result.(*pb.UserLoginResponse)
		testreq.AssertPbEqual(t, &pb.UserLoginResponse{
			User: &pb.User{
				Username: string(testUser.Username),
			},
		}, loginUserResponse)

		exists := db_user.UsernameExists(db, testUser.Username)
		assert.Equal(t, true, exists)
	})
}

func TestInsertUser(t *testing.T) {
	testdb.Open(func(db *sqlx.DB) {
		testUser :=
			db_user.Row{
				Username: "random",
				Password: "WhatAGreatPassword!!1",
			}

		result, err := user.InsertUser(handler.UnauthenticatedParams{
			Message: &pb.UserSignupRequest{
				Username: string(testUser.Username),
				Password: string(testUser.Password),
			},
			DB: db,
		})
		assert.NoError(t, err)

		insertUserResponse := result.(*pb.UserSignupResponse)
		testreq.AssertPbEqual(t, &pb.UserSignupResponse{
			User: &pb.User{
				Username: string(testUser.Username),
			},
		}, insertUserResponse)

		exists := db_user.UsernameExists(db, testUser.Username)
		assert.Equal(t, true, exists)
	})
}

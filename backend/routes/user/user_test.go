package user_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/spencerfcp/scoirapi/v2/database/db_user"
	"github.com/spencerfcp/scoirapi/v2/handler"
	"github.com/spencerfcp/scoirapi/v2/pb"
	"github.com/spencerfcp/scoirapi/v2/routes/user"
	"github.com/spencerfcp/scoirapi/v2/testdb"
	"github.com/spencerfcp/scoirapi/v2/testreq"
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

package user

import (
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/spencerfcp/user-login-code-challenge/backend/database"
	"github.com/spencerfcp/user-login-code-challenge/backend/database/db_user"
	"github.com/spencerfcp/user-login-code-challenge/backend/handler"
	"github.com/spencerfcp/user-login-code-challenge/backend/id"
	"github.com/spencerfcp/user-login-code-challenge/backend/pb"
	"github.com/spencerfcp/user-login-code-challenge/backend/validator"
	"google.golang.org/protobuf/proto"
)

func Login(p handler.UnauthenticatedParams) (proto.Message, error) {
	req := p.Message.(*pb.UserLoginRequest)
	user, err := db_user.GetUserByUsername(p.DB, id.Username(req.Username))
	if err != nil {
		return nil, errors.Wrapf(err, "error fetching user from username %s", req.Username)
	}

	if user == nil || !db_user.PasswordsMatch(user.Password, req.Password) {
		return &pb.UserLoginResponse{
			InvalidCredentials: true,
		}, nil
	}

	return &pb.UserLoginResponse{
		User: &pb.User{
			Username: string(user.Username),
		},
	}, nil
}

func InsertUser(p handler.UnauthenticatedParams) (proto.Message, error) {
	req := p.Message.(*pb.UserSignupRequest)
	safeUsername := id.Username(strings.ToLower(req.Username))

	if len(safeUsername) < 2 || !validator.Password(req.Password) {
		return nil, errors.New("invalid user creation request")
	}

	hashedPassword, err := db_user.HashPassword(req.Password)
	if err != nil {
		return nil, errors.Wrap(err, "unable to hash user password")
	}
	emailExists := db_user.UsernameExists(p.DB, safeUsername)
	if emailExists {
		return &pb.UserSignupResponse{
			UsernameAlreadyExists: true,
		}, nil
	}

	var newUserID id.UserID
	err = database.Transaction(p.DB, func(db *sqlx.Tx) error {
		newUserID, err = db_user.Add(db, db_user.Row{
			Username: safeUsername,
			Password: hashedPassword,
		})

		if err != nil {
			return errors.Wrapf(err, "Unable to insert user")
		}
		return nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "insert user transaction failed")
	}

	newUser, err := db_user.GetUser(p.DB, newUserID)
	if err != nil {
		return nil, err
	}

	return &pb.UserSignupResponse{
		User: &pb.User{
			Username: string(newUser.Username),
		},
	}, nil
}

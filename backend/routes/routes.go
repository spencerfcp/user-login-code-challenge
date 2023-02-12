package routes

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/spencerfcp/user-login-code-challenge/backend/env"
	"github.com/spencerfcp/user-login-code-challenge/backend/handler"
	"github.com/spencerfcp/user-login-code-challenge/backend/pb"
	"github.com/spencerfcp/user-login-code-challenge/backend/routes/user"
	"google.golang.org/protobuf/proto"
)

func Handle(
	mux *http.ServeMux,
	db *sqlx.DB,
	env env.Env,
) {
	handler.Path(mux, "/user", handler.Methods{
		Post: handler.UnauthenticatedProto(
			func() proto.Message { return &pb.UserSignupRequest{} },
			user.InsertUser,
			db,
			env),
	})
	handler.Path(mux, "/login", handler.Methods{
		Post: handler.UnauthenticatedProto(
			func() proto.Message { return &pb.UserLoginRequest{} },
			user.Login,
			db,
			env),
	})

}

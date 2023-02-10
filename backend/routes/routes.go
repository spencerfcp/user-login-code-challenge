package routes

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/spencerfcp/scoirapi/v2/env"
	"github.com/spencerfcp/scoirapi/v2/handler"
	"github.com/spencerfcp/scoirapi/v2/pb"
	"github.com/spencerfcp/scoirapi/v2/routes/user"
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

package gapi

import (
	"context"
	"database/sql"

	"github.com/BerdiyorovAbrorjon/simplebank/pb"
	"github.com/BerdiyorovAbrorjon/simplebank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "error on get user: %s", err)
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "password incorrect")
	}

	token, err := server.tokenMaker.CreateToken(
		req.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error on create token: %s", err)
	}

	res := pb.LoginUserResponse{
		User:        convertUser(user),
		AccessToken: token,
	}

	return &res, nil
}

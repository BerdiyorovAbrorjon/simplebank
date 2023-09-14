package gapi

import (
	"context"

	db "github.com/BerdiyorovAbrorjon/simplebank/db/sqlc"
	"github.com/BerdiyorovAbrorjon/simplebank/pb"
	"github.com/BerdiyorovAbrorjon/simplebank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashedPass, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error on hashing password")
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPass,
		Fullname:       req.GetFullname(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "error on create user: %s", err)
	}

	res := pb.CreateUserResponse{
		User: convertUser(user),
	}

	return &res, nil
}

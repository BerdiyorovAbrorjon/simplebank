package gapi

import (
	db "github.com/BerdiyorovAbrorjon/simplebank/db/sqlc"
	"github.com/BerdiyorovAbrorjon/simplebank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.Users) *pb.User {
	return &pb.User{
		Username:          user.Username,
		Fullname:          user.Fullname,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}

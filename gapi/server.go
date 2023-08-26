package gapi

import (
	"fmt"

	db "github.com/BerdiyorovAbrorjon/simplebank/db/sqlc"
	"github.com/BerdiyorovAbrorjon/simplebank/pb"
	"github.com/BerdiyorovAbrorjon/simplebank/token"
	"github.com/BerdiyorovAbrorjon/simplebank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("create token maker error: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

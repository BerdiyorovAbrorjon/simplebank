package main

import (
	"database/sql"
	"log"

	"github.com/BerdiyorovAbrorjon/simplebank/api"
	db "github.com/BerdiyorovAbrorjon/simplebank/db/sqlc"
	"github.com/BerdiyorovAbrorjon/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("config error", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}

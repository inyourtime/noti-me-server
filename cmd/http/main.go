package main

import (
	"log"

	"github.com/inyourtime/noti-me-server/config"
	"github.com/inyourtime/noti-me-server/internal/adapter/http/server"
	"github.com/inyourtime/noti-me-server/internal/adapter/repository/db"
)

func main() {
	config, err := config.New(config.NewViper())
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.New(config)
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewHttpServer(config, db)
	log.Fatal(s.Start())
}

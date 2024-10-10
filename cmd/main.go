package main

import (
	"log"

	"github.com/LeMinh0706/ChatApp/cmd/server"
	"github.com/LeMinh0706/ChatApp/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	server, err := server.NewServer(config)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}
	server.Start(config.ServerAddress)
}

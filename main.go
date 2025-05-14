package main

import (
	"log"

	"github.com/liquiddev99/mantra-interview-be/api"
	"github.com/liquiddev99/mantra-interview-be/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config file", err)
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("Cannot create server", err)
	}

	err = server.Start("0.0.0.0:9090")

	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}

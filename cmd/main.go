package main

import (
	"log"

	"github.com/marioanchevski/docu-reach/cmd/api"
	"github.com/marioanchevski/docu-reach/config"
)

func main() {

	cfg := config.NewStandardConfig()

	server := api.NewAPIServer(cfg)

	if err := server.Run(); err != nil {
		log.Fatal(err)

	}
}

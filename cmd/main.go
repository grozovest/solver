package main

import (
	"log"

	"solver/internal/app"
	"solver/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	a, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(a.Run())
}

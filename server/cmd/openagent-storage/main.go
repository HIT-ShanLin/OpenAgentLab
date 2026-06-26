package main

import (
	"log"

	"github.com/HIT-ShanLin/OpenAgentLab/server/internal/storage/bootstrap"
)

func main() {
	cfg, err := bootstrap.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	app, err := bootstrap.Wire(cfg)
	if err != nil {
		log.Fatalf("failed to wire dependencies: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("application error: %v", err)
	}
}

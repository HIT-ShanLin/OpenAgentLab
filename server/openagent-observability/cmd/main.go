package main

import (
	"log"

	"github.com/HIT-ShanLin/OpenAgentLab/server/openagent-observability/internal"
)

func main() {
	cfg, err := observability.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	app, err := observability.Wire(cfg)
	if err != nil {
		log.Fatalf("failed to wire dependencies: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("application error: %v", err)
	}
}

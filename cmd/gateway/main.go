package main

import (
	"log"

	"github.com/dinhbaokhanh/Final-Project-API-Gateway/internal/app"
	"github.com/dinhbaokhanh/Final-Project-API-Gateway/internal/config"
)

func main() {
	cfg := config.Load()

	gateway, err := app.New(cfg)
	if err != nil {
		log.Fatalf("failed to initialize gateway: %v", err)
	}

	log.Printf("gateway listening on :%s", cfg.Port)
	if err := gateway.Run(); err != nil {
		log.Fatalf("gateway stopped with error: %v", err)
	}
}

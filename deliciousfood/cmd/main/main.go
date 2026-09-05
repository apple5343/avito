package main

import (
	"deliciousfood/internal/config"
	"deliciousfood/internal/transport"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	server := transport.NewServer(cfg)

	log.Printf("Delicious Food restaurant service starting on %s", cfg.Addr)
	log.Printf("Connected to Avito at %s", cfg.AvitoURL)

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

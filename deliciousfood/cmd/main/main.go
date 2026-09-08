package main

import (
	"deliciousfood/internal/config"
	"deliciousfood/internal/register"
	"deliciousfood/internal/transport"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	callbackURL := cfg.CallbackURL
	if callbackURL == "" {
		callbackURL = fmt.Sprintf("http://localhost%s/orders", cfg.Addr)
	}

	if cfg.APIKey == "" {
		apiKey, err := register.EnsureRegistered(cfg.AvitoURL, register.Restaurant{
			Name:        cfg.Name,
			Description: "Вкусная еда с доставкой",
			City:        cfg.City,
			Cuisine:     cfg.Cuisine,
			CallbackURL: callbackURL,
			IsActive:    true,
		})
		if err != nil {
			log.Fatalf("Failed to register restaurant: %v", err)
		}
		cfg.APIKey = apiKey
	}

	server := transport.NewServer(cfg)

	log.Printf("Delicious Food restaurant service starting on %s", cfg.Addr)
	log.Printf("Connected to Avito at %s", cfg.AvitoURL)

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

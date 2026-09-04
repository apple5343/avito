package main

import (
	"avito/internal/config"
	"avito/internal/infrastructure"
	restaurantRepository "avito/internal/repository/restaurant"
	restaurantService "avito/internal/service/restaurant"
	"avito/internal/transport"
)

func main() {
	dbConfig, err := config.PostgresConfig()
	if err != nil {
		panic(err)
	}

	serverConfig, err := config.ServerConfig()
	if err != nil {
		panic(err)
	}

	db, err := infrastructure.NewPostgres(dbConfig)
	if err != nil {
		panic(err)
	}

	restaurantRepo := restaurantRepository.NewRepository(db)
	restaurantService := restaurantService.NewService(restaurantRepo)

	server := transport.NewServer(serverConfig, restaurantService)
	if err := server.Start(); err != nil {
		panic(err)
	}
}

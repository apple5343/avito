package main

import (
	"avito/internal/config"
	"avito/internal/infrastructure"
	menuRepository "avito/internal/repository/menu"
	restaurantRepository "avito/internal/repository/restaurant"
	menuService "avito/internal/service/menu"
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
	menuRepo := menuRepository.NewRepository(db)

	restaurantService := restaurantService.NewService(restaurantRepo)
	menuService := menuService.NewService(menuRepo, restaurantService)

	server := transport.NewServer(serverConfig, restaurantService, menuService)
	err = server.Start()
	if err != nil {
		panic(err)
	}
}

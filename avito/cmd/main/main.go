package main

import (
	"avito/internal/config"
	"avito/internal/infrastructure"
	"avito/internal/pkg/restaurantapi/http"
	menuRepository "avito/internal/repository/menu"
	orderRepository "avito/internal/repository/order"
	restaurantRepository "avito/internal/repository/restaurant"
	menuService "avito/internal/service/menu"
	orderService "avito/internal/service/order"
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
	orderRepo := orderRepository.NewRepository(db)

	restaurantService := restaurantService.NewService(restaurantRepo)
	menuService := menuService.NewService(menuRepo, restaurantService)
	orderService := orderService.NewService(orderRepo, menuService, http.NewClient(), restaurantService)

	server := transport.NewServer(serverConfig, restaurantService, menuService, orderService)
	err = server.Start()
	if err != nil {
		panic(err)
	}
}

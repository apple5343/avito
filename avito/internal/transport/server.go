package transport

import (
	"avito/internal/config"
	"avito/internal/transport/menu"
	"avito/internal/transport/middleware"
	"avito/internal/transport/restaurant"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e                 *echo.Echo
	restaurantHandler *restaurant.Handler
	menuHandler       *menu.Handler
	cfg               *config.Server
}

func NewServer(cfg *config.Server, restaurantService restaurant.Service, menuService menu.Service) *Server {
	return &Server{
		e:                 echo.New(),
		restaurantHandler: restaurant.NewHandler(restaurantService),
		menuHandler:       menu.NewHandler(menuService),
		cfg:               cfg,
	}
}

func (s *Server) Start() error {
	s.route()
	return s.e.Start(s.cfg.Addr)
}

func (s *Server) route() {
	s.e.Use(middleware.ErrorMiddleware())
	restaurant := s.e.Group("/restaurants")
	{
		restaurant.GET("", s.restaurantHandler.List())
		restaurant.POST("", s.restaurantHandler.Create())
		restaurant.PATCH("", middleware.APIKeyMiddleware()(s.restaurantHandler.Update()))
		restaurant.GET("/:id", s.restaurantHandler.GetByID())
		restaurant.GET("/:id/menu", s.menuHandler.GetMenu())
	}
	menu := s.e.Group("/menu")
	{
		menu.POST("", middleware.APIKeyMiddleware()(s.menuHandler.Create()))
		menu.PATCH("/:id", middleware.APIKeyMiddleware()(s.menuHandler.Update()))
	}
}

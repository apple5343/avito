package transport

import (
	"avito/internal/config"
	"avito/internal/transport/middleware"
	"avito/internal/transport/restaurant"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e                 *echo.Echo
	restaurantHandler *restaurant.Handler
	cfg               *config.Server
}

func NewServer(cfg *config.Server, restaurantService restaurant.Service) *Server {
	return &Server{
		e:                 echo.New(),
		restaurantHandler: restaurant.NewHandler(restaurantService),
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
		restaurant.POST("", s.restaurantHandler.Create())
		restaurant.PATCH("", middleware.APIKeyMiddleware()(s.restaurantHandler.Update()))
	}
}

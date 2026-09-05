package transport

import (
	"deliciousfood/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	e   *echo.Echo
	cfg *config.Server
}

func NewServer(cfg *config.Server) *Server {
	return &Server{
		e:   echo.New(),
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	s.route()
	return s.e.Start(s.cfg.Addr)
}

func (s *Server) route() {
	s.e.Use(middleware.Recover())

	s.e.POST("/orders", s.OrderHandler())
}

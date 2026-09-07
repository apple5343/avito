package transport

import (
	"net/http"

	"avito/internal/config"
	"avito/internal/transport/menu"
	"avito/internal/transport/middleware"
	"avito/internal/transport/order"
	"avito/internal/transport/restaurant"

	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e                 *echo.Echo
	restaurantHandler *restaurant.Handler
	menuHandler       *menu.Handler
	orderHandler      *order.Handler
	cfg               *config.Server
}

func NewServer(cfg *config.Server, restaurantService restaurant.Service, menuService menu.Service, orderService order.Service) *Server {
	return &Server{
		e:                 echo.New(),
		restaurantHandler: restaurant.NewHandler(restaurantService),
		menuHandler:       menu.NewHandler(menuService),
		orderHandler:      order.NewHandler(orderService),
		cfg:               cfg,
	}
}

func (s *Server) Start() error {
	s.route()
	return s.e.Start(s.cfg.Addr)
}

func (s *Server) route() {
	s.e.Use(echoMiddleware.Recover())
	s.e.Use(middleware.ErrorMiddleware())
	s.e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})

	})

	restaurant := s.e.Group("/restaurants")
	{
		restaurant.GET("", s.restaurantHandler.List())
		restaurant.POST("", s.restaurantHandler.Create())
		restaurant.PATCH("/:id", middleware.APIKeyMiddleware()(s.restaurantHandler.Update()))
		restaurant.GET("/:id", s.restaurantHandler.GetByID())
		restaurant.GET("/:id/menu", s.menuHandler.GetMenu())
	}
	menu := s.e.Group("/menu")
	{
		menu.POST("", middleware.APIKeyMiddleware()(s.menuHandler.Create()))
		menu.PATCH("/:id", middleware.APIKeyMiddleware()(s.menuHandler.Update()))
	}
	order := s.e.Group("/orders")
	{
		order.GET("", middleware.UserIDMiddleware()(s.orderHandler.GetUserOrders()))
		order.GET("/:id", middleware.UserIDMiddleware()(s.orderHandler.GetUserOrder()))
		order.PATCH("/:id/status", middleware.APIKeyMiddleware()(s.orderHandler.UpdateStatus()))
		order.GET("/:id/history", middleware.UserIDMiddleware()(s.orderHandler.GetOrderHistory()))
		order.POST("", middleware.UserIDMiddleware()(s.orderHandler.Create()))
	}
}

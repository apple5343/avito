package transport

import (
	"bytes"
	"context"
	"deliciousfood/internal/model"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var issues = []string{
	"В связи с повышенной нагрузкой, доставка в данном районе приостановлена.",
	"К сожалению, в данный момент нет свободных курьеров.",
	"Некоторые позиции из заказа недоступны.",
}

func (s *Server) OrderHandler() echo.HandlerFunc {
	type request struct {
		Order model.Order       `json:"order"`
		Items []model.OrderItem `json:"items"`
	}
	type response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	return func(c echo.Context) error {
		var req request
		if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response{
				Status:  "error",
				Message: "invalid json",
			})
		}

		if rand.IntN(100) < 50 { //nolint
			return c.JSON(http.StatusConflict, response{
				Status:  "error",
				Message: issues[rand.IntN(len(issues))], //nolint
			})
		}

		go s.processOrder(req.Order.ID, s.cfg.AvitoURL, s.cfg.APIKey)

		return c.JSON(http.StatusCreated, response{
			Status:  "success",
			Message: "order created",
		})
	}
}

func (s *Server) processOrder(orderID int, avitoURL, apiKey string) {
	statusFlow := []struct {
		status  string
		message string
	}{
		{"cooking", "Ваш заказ готовится"},
		{"ready", "Заказ готов, ожидает курьера"},
		{"delivering", "Курьер выехал к вам"},
		{"delivered", "Заказ доставлен. Приятного аппетита!"},
	}

	for _, step := range statusFlow {
		sleepDuration := time.Second * time.Duration(rand.IntN(20)+10) //nolint
		log.Printf("Order %d: waiting %v before status '%s'", orderID, sleepDuration, step.status)
		time.Sleep(sleepDuration)

		if err := s.sendStatusUpdate(orderID, step.status, step.message, avitoURL, apiKey); err != nil {
			log.Printf("Failed to update order %d status to %s: %v", orderID, step.status, err)
			return
		}

		log.Printf("Order %d: status updated to '%s'", orderID, step.status)
	}
}

func (s *Server) sendStatusUpdate(orderID int, status, message, avitoURL, apiKey string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer cancel()
	type updateRequest struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	reqBody := updateRequest{Status: status, Message: message}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	url := fmt.Sprintf("%s/orders/%d/status", avitoURL, orderID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Printf("close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

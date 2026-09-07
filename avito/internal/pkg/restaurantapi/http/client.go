package http

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"avito/internal/model"
)

const retryCount = 3
const requestTimeout = time.Second * 2

type Client struct {
	c *http.Client
}

func NewClient() *Client {
	return &Client{
		c: &http.Client{
			Timeout: requestTimeout,
		},
	}
}

type request struct {
	Order *model.Order      `json:"order"`
	Items []model.OrderItem `json:"items"`
}

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (c *Client) Create(ctx context.Context, url string, order *model.Order, items []model.OrderItem) (bool, string, error) {
	body, err := json.Marshal(request{Order: order, Items: items})
	if err != nil {
		return false, "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return false, "", err
	}
	req.Header.Set("Content-Type", "application/json")

	timeout := time.Second
	var resp *http.Response
	for i := 0; i < retryCount; i++ {
		resp, err = c.c.Do(req)
		if err == nil {
			break
		}
		time.Sleep(timeout)
		timeout *= 2
	}
	if err != nil {
		return false, "", err
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Printf("close response body: %v", err)
		}
	}()

	if resp.StatusCode == http.StatusCreated {
		var res response
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return false, "failed to parse restaurant response", nil
		}
		return true, res.Message, nil
	}

	if resp.StatusCode == http.StatusConflict {
		var res response
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return false, "restaurant rejected the order", nil
		}
		return false, res.Message, nil
	}

	return false, "restaurant returned unexpected status", nil
}

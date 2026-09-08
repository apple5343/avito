package register

import (
	"bytes"
	"context"
	"deliciousfood/internal/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	apiKeyFile     = ".api_key"
	maxRetries     = 10
	retryInterval  = 3 * time.Second
	requestTimeout = 10 * time.Second
)

type Restaurant struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	City        string `json:"city"`
	Cuisine     string `json:"cuisine"`
	CallbackURL string `json:"callback_url"`
	IsActive    bool   `json:"is_active"`
}

type RegisterResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	City        string `json:"city"`
	Cuisine     string `json:"cuisine"`
	CallbackURL string `json:"callback_url"`
	APIKey      string `json:"api_key"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func EnsureRegistered(avitoURL string, restaurant Restaurant) (string, error) {
	if apiKey, err := loadAPIKey(); err == nil && apiKey != "" {
		log.Printf("Api key found")
		log.Printf("API key: %s", apiKey)
		return apiKey, nil
	}

	log.Printf("API key not found, registering restaurant...")

	apiKey, err := registerRestaurant(avitoURL, restaurant)
	if err != nil {
		return "", fmt.Errorf("failed to register restaurant: %w", err)
	}

	if err := saveAPIKey(apiKey); err != nil {
		return "", fmt.Errorf("failed to save API key: %w", err)
	}

	log.Printf("API key: %s", apiKey)
	log.Printf("Restaurant registered, API key saved to %s", apiKeyFile)
	log.Printf("Add menu items...")
	if err := addMenuItems(avitoURL, apiKey, model.MenuItems); err != nil {
		return "", fmt.Errorf("failed to add menu items: %w", err)
	}
	log.Printf("Menu added")
	return apiKey, nil
}

func registerRestaurant(avitoURL string, restaurant Restaurant) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	url := avitoURL + "/restaurants"

	jsonData, err := json.Marshal(restaurant)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: requestTimeout}
	var resp *http.Response
	for i := 0; i < maxRetries; i++ {
		resp, err = client.Do(req)
		if err == nil {
			break
		}
		time.Sleep(retryInterval)
	}

	if err != nil {
		return "", fmt.Errorf("send request: %w", err)
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Printf("close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result RegisterResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	log.Printf("Restaurant registered: ID=%d, Name=%s", result.ID, result.Name)
	return result.APIKey, nil
}

func loadAPIKey() (string, error) {
	data, err := os.ReadFile(apiKeyFile)
	if err != nil {
		return "", err
	}
	return string(bytes.TrimSpace(data)), nil
}

func saveAPIKey(apiKey string) error {
	return os.WriteFile(apiKeyFile, []byte(apiKey), 0600)
}

func addMenuItems(avitoURL string, apiKey string, menuItems []model.MenuItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for _, menuItem := range menuItems {
		url := avitoURL + "/menu"
		jsonData, err := json.Marshal(menuItem)
		if err != nil {
			return fmt.Errorf("marshal request: %w", err)
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonData))
		if err != nil {
			return fmt.Errorf("create request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-API-Key", apiKey)

		client := &http.Client{Timeout: requestTimeout}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("send request: %w", err)
		}
		defer func() {
			if err = resp.Body.Close(); err != nil {
				log.Printf("close response body: %v", err)
			}
		}()
		if resp.StatusCode != http.StatusCreated {
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
	}
	return nil
}

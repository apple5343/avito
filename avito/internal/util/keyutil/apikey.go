package keyutil

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateAPIKey() (string, error) {
	const apiKeyLength = 32
	bytes := make([]byte, apiKeyLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate API key: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}

package keyutil

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateAPIKey() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate API key: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}

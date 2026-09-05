package model

import (
	"errors"
	"fmt"
	"time"

	"avito/internal/util/keyutil"
	"avito/pkg/errorx"
	"avito/pkg/validator"
)

var Cuisines = map[string]bool{
	"japanese": true,
	"chinese":  true,
	"italian":  true,
	"russian":  true,
	"american": true,
	"european": true,
	"other":    true,
}

type Restaurant struct {
	ID          int       `json:"id"           db:"id"`
	Name        string    `json:"name"         db:"name"`
	Description string    `json:"description"  db:"description"`
	City        string    `json:"city"         db:"city"`
	Cuisine     string    `json:"cuisine"      db:"cuisine"`
	CallbackURL string    `json:"callback_url" db:"callback_url"`
	APIKey      string    `json:"api_key"      db:"api_key"`
	IsActive    bool      `json:"is_active"    db:"is_active"`
	CreatedAt   time.Time `json:"created_at"   db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"   db:"updated_at"`
}

func (r *Restaurant) BeforeCreate() error {
	if len(r.Name) == 0 {
		return errorx.New(errorx.CodeBadRequest, errors.New("name is required"))
	}
	if len(r.City) == 0 {
		return errorx.New(errorx.CodeBadRequest, errors.New("city is required"))
	}
	if len(r.Cuisine) == 0 {
		return errorx.New(errorx.CodeBadRequest, errors.New("cuisine is required"))
	}
	if !Cuisines[r.Cuisine] {
		return errorx.New(errorx.CodeBadRequest, errors.New("invalid cuisine type"))
	}
	if isValid := validator.ValidateURL(r.CallbackURL); !isValid {
		return errorx.New(errorx.CodeBadRequest, errors.New("invalid callback url"))
	}

	apiKey, err := keyutil.GenerateAPIKey()
	if err != nil {
		return errorx.New(errorx.CodeInternalServerError, fmt.Errorf("generate API key: %w", err))
	}

	r.APIKey = apiKey
	return nil
}

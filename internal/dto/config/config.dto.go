package dto

import (
	"time"

	"github.com/google/uuid"
)

// Config DTOs

type ConfigCreateRequest struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
}

type ConfigUpdateRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type ConfigResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

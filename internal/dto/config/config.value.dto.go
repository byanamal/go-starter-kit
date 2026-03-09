package dto

import (
	"time"

	"github.com/google/uuid"
)

type ConfigValueCreateRequest struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
}

type ConfigValueUpdateRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type ConfigValueResponse struct {
	ID        uuid.UUID `json:"id"`
	ConfigID  uuid.UUID `json:"config_id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

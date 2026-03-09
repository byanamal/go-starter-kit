package dto

import (
	"time"

	"github.com/google/uuid"
)

type PermissionResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Group     *string   `json:"group"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GroupedPermissionResponse struct {
	Group string               `json:"group"`
	Items []PermissionResponse `json:"items"`
}

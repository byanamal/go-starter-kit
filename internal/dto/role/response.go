package dto

import (
	"time"

	"github.com/google/uuid"
)

type RoleResponse struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Code        string      `json:"code"`
	Description *string     `json:"description"`
	CreatedAt   *time.Time  `json:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at"`
	Permissions interface{} `json:"permissions"`
}

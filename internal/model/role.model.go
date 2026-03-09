package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID       `db:"id" json:"id"`
	Name        string          `db:"name" json:"name"`
	Code        string          `db:"code" json:"code"`
	Description *string         `db:"description" json:"description"`
	CreatedAt   *time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time      `db:"updated_at" json:"updated_at"`
	Permissions json.RawMessage `db:"permissions" json:"permissions"`
}

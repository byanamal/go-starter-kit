package model

import (
	"time"

	"github.com/google/uuid"
)

type ConfigValues struct {
	ID        uuid.UUID `db:"id" json:"id"`
	ConfigID  uuid.UUID `db:"config_id" json:"config_id"`
	Name      string    `db:"name" json:"name"`
	Code      string    `db:"code" json:"code"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

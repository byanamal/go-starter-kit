package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Email    string    `db:"email" json:"email"`
	Name     *string   `db:"name" json:"name"`
	Password string    `db:"password" json:"password"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`


	Roles   json.RawMessage `db:"roles"`
}

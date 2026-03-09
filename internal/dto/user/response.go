package dto

import (

	dtoRole "base-api/internal/dto/role"
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  *string   `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`


	Roles   []dtoRole.RoleResponse         `json:"roles"`
}

type UserWidgetResponse struct {
	TotalUsers    int64 `db:"total_users" json:"total_users"`
	ActiveUsers   int64 `db:"active_users" json:"active_users"`
	InactiveUsers int64 `db:"inactive_users" json:"inactive_users"`

	ThisMonthUsers int64 `db:"this_month_users" json:"this_month_users"`
	LastMonthUsers int64 `db:"last_month_users" json:"last_month_users"`

	ThisWeekUsers int64 `db:"this_week_users" json:"this_week_users"`
	LastWeekUsers int64 `db:"last_week_users" json:"last_week_users"`
}

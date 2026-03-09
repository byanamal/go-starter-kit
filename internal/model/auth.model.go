package model

import "github.com/google/uuid"

type AuthUser struct {
	ID          uuid.UUID `db:"id"`
	Email       string    `db:"email"`
	Password    string    `db:"password" json:"password"`
	Roles       []string
	Permissions []string
}

type UserWithRole struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Email    string    `db:"email" json:"email"`
	Name     *string   `db:"name" json:"name"`
	Password string    `db:"password" json:"password"`
	RoleName string    `db:"role_name" json:"role_name"`
	RoleSlug string    `db:"role_slug" json:"role_slug"`
}

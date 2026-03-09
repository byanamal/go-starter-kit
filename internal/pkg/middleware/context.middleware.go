package middleware

import (
	"context"

	"github.com/google/uuid"
)

type contextKey string

const (
	ctxUserID      contextKey = "user_id"
	ctxEmail       contextKey = "email"
	ctxRoles       contextKey = "roles"
	ctxPermissions contextKey = "permissions"
)

func GetUserID(ctx context.Context) (uuid.UUID, bool) {
	userID, ok := ctx.Value(ctxUserID).(uuid.UUID)
	return userID, ok
}

func GetEmail(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(ctxEmail).(string)
	return email, ok
}

func GetRoles(ctx context.Context) ([]string, bool) {
	roles, ok := ctx.Value(ctxRoles).([]string)
	return roles, ok
}

func GetPermissions(ctx context.Context) ([]string, bool) {
	roles, ok := ctx.Value(ctxPermissions).([]string)
	return roles, ok
}

package middleware

import (
	"context"
	"net/http"
	"strings"

	"base-api/internal/pkg/constants"
	"base-api/internal/pkg/helper"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helper.WriteError(w, http.StatusUnauthorized, constants.UNAUTHORIZED)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			helper.WriteError(w, http.StatusUnauthorized, constants.UNAUTHORIZED)
			return
		}

		claims, err := helper.ValidateToken(parts[1])
		if err != nil {
			helper.WriteError(w, http.StatusUnauthorized, constants.UNAUTHORIZED)
			return
		}

		ctx := r.Context()

		ctx = context.WithValue(ctx, ctxUserID, claims.UserID)
		ctx = context.WithValue(ctx, ctxEmail, claims.Email)
		ctx = context.WithValue(ctx, ctxRoles, claims.Roles)
		ctx = context.WithValue(ctx, ctxPermissions, claims.Permissions)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Permission(required ...string) func(http.Handler) http.Handler {
	requiredSet := make(map[string]struct{}, len(required))
	for _, p := range required {
		requiredSet[p] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			perms, ok := r.Context().Value(ctxPermissions).([]string)
			if !ok {
				helper.WriteError(w, http.StatusForbidden, constants.FORBIDDEN)
				return
			}

			userPerms := make(map[string]struct{}, len(perms))
			for _, p := range perms {
				userPerms[p] = struct{}{}
			}

			for p := range requiredSet {
				if _, ok := userPerms[p]; !ok {
					helper.WriteError(w, http.StatusForbidden, constants.FORBIDDEN)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

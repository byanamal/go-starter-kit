package seeder

import (
	"base-api/internal/model"
	"base-api/internal/pkg/constants"
	"base-api/internal/pkg/helper"
	"base-api/internal/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func seedPermissions(ctx context.Context, tx *sqlx.Tx, db *sqlx.DB) error {
	permRepo := repository.NewPermissionRepository(db)

	var adminRoleID, userRoleID uuid.UUID
	err := tx.GetContext(ctx, &adminRoleID, "SELECT id FROM roles WHERE LOWER(name) = $1", "admin")
	if err != nil {
		return fmt.Errorf("failed to get admin role: %w", err)
	}
	err = tx.GetContext(ctx, &userRoleID, "SELECT id FROM roles WHERE LOWER(name) = $1", "user")
	if err != nil {
		return fmt.Errorf("failed to get user role: %w", err)
	}

	for _, module := range constants.Modules {
		for _, action := range constants.Actions {
			code := constants.GetPermissionCode(module, action)
			name := constants.GetPermissionName(module, action)

			permID, err := permRepo.UpsertPermission(ctx, tx, model.Permission{
				Name:  name,
				Code:  code,
				Group: helper.String(constants.TitleCase(module)),
			})
			if err != nil {
				return fmt.Errorf("failed to upsert permission %s: %w", name, err)
			}

			if err := permRepo.AssignToRole(ctx, tx, adminRoleID, permID); err != nil {
				return fmt.Errorf("failed to assign permission %s to admin: %w", name, err)
			}

			if action == constants.ActionView {
				if err := permRepo.AssignToRole(ctx, tx, userRoleID, permID); err != nil {
					return fmt.Errorf("failed to assign permission %s to user: %w", name, err)
				}
			}
		}
	}

	return nil
}

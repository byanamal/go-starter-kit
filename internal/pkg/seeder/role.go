package seeder

import (
	"base-api/internal/model"
	"base-api/internal/pkg/constants"
	"base-api/internal/pkg/helper"
	"base-api/internal/repository"
	"context"

	"github.com/jmoiron/sqlx"
)

func seedRoles(ctx context.Context, tx *sqlx.Tx, db *sqlx.DB) error {
	roleRepo := repository.NewRoleRepository(db)

	roles := []model.Role{
		{
			Name:        "Super Admin",
			Code:        constants.RoleCodeSuperadmin,
			Description: helper.String("Super Administrator with absolute access"),
		},
		{
			Name:        "Admin",
			Code:        constants.RoleCodeAdmin,
			Description: helper.String("Administrator with full access"),
		},
		{
			Name:        "User",
			Code:        constants.RoleCodeUser,
			Description: helper.String("Regular User"),
		},
	}

	for _, role := range roles {
		if _, err := roleRepo.UpsertRole(ctx, tx, role); err != nil {
			return err
		}
	}

	return nil
}

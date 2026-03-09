package seeder

import (
	"base-api/internal/model"
	"base-api/internal/pkg/constants"
	"base-api/internal/pkg/helper"
	"base-api/internal/repository"
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func seedUser(ctx context.Context, tx *sqlx.Tx, db *sqlx.DB) error {
	userRepo := repository.NewUserRepository(db)

	password, err := helper.HashPassword("admin")
	if err != nil {
		return err
	}

	userID, err := userRepo.UpsertUser(ctx, tx, model.User{
		Name:     helper.String("admin@admin.com"),
		Email:    "admin@admin.com",
		Password: password,
	})
	if err != nil {
		return err
	}

	// Get admin role
	var adminRoleID uuid.UUID
	err = tx.GetContext(ctx, &adminRoleID, "SELECT id FROM roles WHERE code = $1", constants.RoleCodeAdmin)
	if err != nil {
		return err
	}

	// Assign role
	err = userRepo.AssignRole(ctx, tx, userID, adminRoleID)
	if err != nil {
		return err
	}

	return nil
}

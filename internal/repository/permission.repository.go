package repository

import (
	"base-api/internal/model"
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PermissionRepository interface {
	FindAll(ctx context.Context) ([]model.Permission, error)
	FindByID(ctx context.Context, id uuid.UUID) (model.Permission, error)
	UpsertPermission(ctx context.Context, tx *sqlx.Tx, req model.Permission) (uuid.UUID, error)
	AssignToRole(ctx context.Context, tx *sqlx.Tx, roleID, permissionID uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type permissionRepository struct {
	db *sqlx.DB
}

func NewPermissionRepository(db *sqlx.DB) PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) FindAll(ctx context.Context) ([]model.Permission, error) {
	var permissions []model.Permission

	query := `
		SELECT id, name, code, "group", description, created_at, updated_at
		FROM permissions
		ORDER BY created_at DESC
	`

	err := r.db.SelectContext(ctx, &permissions, query)
	return permissions, err
}

func (r *permissionRepository) FindByID(ctx context.Context, id uuid.UUID) (model.Permission, error) {
	var permission model.Permission

	query := `
		SELECT id, name, code, "group", description, created_at, updated_at
		FROM permissions
		WHERE id = $1
	`

	err := r.db.GetContext(ctx, &permission, query, id)
	return permission, err
}

func (r *permissionRepository) UpsertPermission(ctx context.Context, tx *sqlx.Tx, req model.Permission) (uuid.UUID, error) {
	query := `
		INSERT INTO permissions (name, code, "group", description)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (code) DO UPDATE SET
			name = EXCLUDED.name,
			"group" = EXCLUDED."group",
			description = EXCLUDED.description,
			updated_at = NOW()
		RETURNING id
	`
	var id uuid.UUID
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, req.Name, req.Code, req.Group, req.Description).Scan(&id)
	} else {
		err = r.db.QueryRowContext(ctx, query, req.Name, req.Code, req.Group, req.Description).Scan(&id)
	}
	return id, err
}

func (r *permissionRepository) AssignToRole(ctx context.Context, tx *sqlx.Tx, roleID, permissionID uuid.UUID) error {
	query := `
		INSERT INTO role_permissions (role_id, permission_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, roleID, permissionID)
	} else {
		_, err = r.db.ExecContext(ctx, query, roleID, permissionID)
	}
	return err
}
func (r *permissionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM permissions WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

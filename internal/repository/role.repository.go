package repository

import (
	"base-api/internal/model"
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RoleRepository interface {
	FindAll(ctx context.Context) ([]model.Role, error)
	FindByID(ctx context.Context, id uuid.UUID) (model.Role, error)
	UpsertRole(ctx context.Context, tx *sqlx.Tx, req model.Role) (uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type roleRepository struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) FindAll(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role

	query := `
		SELECT
			r.id,
			r.name,
			r.code,
			r.description,
			r.created_at,
			r.updated_at,
			COALESCE(
				json_agg(
					DISTINCT jsonb_build_object(
						'id', p.id,
						'name', p.name,
						'code', p.code,
						'group', p.group
					)
				) FILTER (WHERE p.id IS NOT NULL),
				'[]'
			) AS permissions
		FROM roles r
		LEFT JOIN role_permissions rp ON rp.role_id = r.id
		LEFT JOIN permissions p ON p.id = rp.permission_id
		GROUP BY r.id
		ORDER BY r.created_at DESC
	`

	err := r.db.SelectContext(ctx, &roles, query)
	return roles, err
}

func (r *roleRepository) FindByID(ctx context.Context, id uuid.UUID) (model.Role, error) {
	var role model.Role

	query := `
		SELECT id, name, code, description, created_at, updated_at
		FROM roles
		WHERE id = $1
	`

	err := r.db.GetContext(ctx, &role, query, id)
	return role, err
}

func (r *roleRepository) UpsertRole(ctx context.Context, tx *sqlx.Tx, req model.Role) (uuid.UUID, error) {
	query := `
		INSERT INTO roles (name, code, description)
		VALUES ($1, $2, $3)
		ON CONFLICT (code) DO UPDATE SET
			name = EXCLUDED.name,
			description = EXCLUDED.description,
			updated_at = NOW()
		RETURNING id
	`
	var id uuid.UUID
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, req.Name, req.Code, req.Description).Scan(&id)
	} else {
		err = r.db.QueryRowContext(ctx, query, req.Name, req.Code, req.Description).Scan(&id)
	}
	return id, err
}
func (r *roleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM roles WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

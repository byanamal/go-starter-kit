package repository

import (
	"base-api/internal/model"
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type UserRepository interface {
	FindAll(ctx context.Context, limit, offset int) ([]model.User, error)
	Count(ctx context.Context) (int, error)
	FindByID(ctx context.Context, id uuid.UUID) (model.User, error)
	FindAuthUserByEmail(
		ctx context.Context,
		email string,
	) (*model.AuthUser, error)
	UpsertUser(ctx context.Context, tx *sqlx.Tx, req model.User) (uuid.UUID, error)
	AssignRole(ctx context.Context, tx *sqlx.Tx, userID, roleID uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{db: db}
}


func (r *userRepository) FindAll(ctx context.Context, limit, offset int) ([]model.User, error) {
	var users []model.User

	query := `
		SELECT
			u.id,
			u.name,
			u.email,
			u.created_at,
			u.updated_at,

			COALESCE(
				json_agg(
					DISTINCT jsonb_build_object(
						'id', r.id,
						'name', r.name,
						'code', r.code
					)
				) FILTER (WHERE r.id IS NOT NULL),
				'[]'
			) AS roles

		FROM users u
		LEFT JOIN user_roles ur ON ur.user_id = u.id
		LEFT JOIN roles r ON r.id = ur.role_id

		GROUP BY
			u.id

		ORDER BY u.created_at DESC
		LIMIT $1 OFFSET $2;
	`

	err := r.db.SelectContext(ctx, &users, query, limit, offset)
	return users, err
}

func (r *userRepository) Count(ctx context.Context) (int, error) {
	var total int

	query := `
		SELECT COUNT(*) FROM users
	`

	err := r.db.GetContext(ctx, &total, query)
	return total, err
}

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (model.User, error) {
	var user model.User

	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
		WHERE id = $1;
	`

	err := r.db.GetContext(ctx, &user, query, id)
	return user, err
}

func (r *userRepository) FindAuthUserByEmail(
	ctx context.Context,
	email string,
) (*model.AuthUser, error) {

	query := `
		SELECT
		    u.id,
		    u.email,
		    u.password,
		    ARRAY_AGG(DISTINCT r.code)::text[] AS roles,
    		ARRAY_AGG(DISTINCT p.code)::text[] AS permissions
		FROM users u
		JOIN user_roles ur ON ur.user_id = u.id
		JOIN roles r ON r.id = ur.role_id
		JOIN role_permissions rp ON rp.role_id = r.id
		JOIN permissions p ON p.id = rp.permission_id
		WHERE u.email = $1
		GROUP BY u.id, u.email, u.password;
	`

	var user model.AuthUser

	err := r.db.QueryRowContext(ctx, query, email).
		Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			pq.Array(&user.Roles),
			pq.Array(&user.Permissions),
		)

	return &user, err
}

func (r *userRepository) UpsertUser(ctx context.Context, tx *sqlx.Tx, req model.User) (uuid.UUID, error) {
	query := `
		INSERT INTO users (email, name, password)
		VALUES ($1, $2, $3)
		ON CONFLICT (email) DO UPDATE SET
			name = EXCLUDED.name,
			password = EXCLUDED.password,
			updated_at = NOW()
		RETURNING id
	`
	var id uuid.UUID
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, req.Email, req.Name, req.Password).Scan(&id)
	} else {
		err = r.db.QueryRowContext(ctx, query, req.Email, req.Name, req.Password).Scan(&id)
	}
	return id, err
}

func (r *userRepository) AssignRole(ctx context.Context, tx *sqlx.Tx, userID, roleID uuid.UUID) error {
	query := `
		INSERT INTO user_roles (user_id, role_id)
		VALUES ($1, $2)
		ON CONFLICT (user_id, role_id) DO NOTHING
	`
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, userID, roleID)
	} else {
		_, err = r.db.ExecContext(ctx, query, userID, roleID)
	}
	return err
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

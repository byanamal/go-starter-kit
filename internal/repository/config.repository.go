package repository

import (
	"base-api/internal/model"
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ConfigRepository interface {
	// Config CRUD
	FindAll(ctx context.Context) ([]model.Config, error)
	FindByID(ctx context.Context, id uuid.UUID) (model.Config, error)
	UpsertConfig(ctx context.Context, tx *sqlx.Tx, req model.Config) (uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type configRepository struct {
	db *sqlx.DB
}

func NewConfigRepository(db *sqlx.DB) ConfigRepository {
	return &configRepository{db: db}
}

// Config Implementation

func (r *configRepository) FindAll(ctx context.Context) ([]model.Config, error) {
	var configs []model.Config
	query := `SELECT id, name, code, created_at, updated_at FROM config ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &configs, query)
	return configs, err
}

func (r *configRepository) FindByID(ctx context.Context, id uuid.UUID) (model.Config, error) {
	var cfg model.Config
	query := `SELECT id, name, code, created_at, updated_at FROM config WHERE id = $1`
	err := r.db.GetContext(ctx, &cfg, query, id)
	return cfg, err
}

func (r *configRepository) UpsertConfig(ctx context.Context, tx *sqlx.Tx, req model.Config) (uuid.UUID, error) {
	query := `
		INSERT INTO config (name, code)
		VALUES ($1, $2)
		ON CONFLICT (code)
		DO UPDATE SET
			name = EXCLUDED.name,
			code = EXCLUDED.code,
			updated_at = NOW()
		RETURNING id
	`
	var id uuid.UUID
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, query, req.Name, req.Code).Scan(&id)
	} else {
		err = r.db.QueryRowContext(ctx, query, req.Name, req.Code).Scan(&id)
	}
	return id, err
}
func (r *configRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM config WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

package repository

import (
	"base-api/internal/model"
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ConfigValueRepository interface {
	FindValuesByConfigID(ctx context.Context, configID uuid.UUID) ([]model.ConfigValues, error)
	FindByID(ctx context.Context, id uuid.UUID) (model.ConfigValues, error)
	UpsertConfigValue(ctx context.Context, tx *sqlx.Tx, req model.ConfigValues) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type configValueRepository struct {
	db *sqlx.DB
}

func NewConfigValueRepository(db *sqlx.DB) ConfigValueRepository {
	return &configValueRepository{db: db}
}

func (r *configValueRepository) FindValuesByConfigID(ctx context.Context, configID uuid.UUID) ([]model.ConfigValues, error) {
	var values []model.ConfigValues
	query := `SELECT id, config_id, name, code, created_at, updated_at FROM config_values WHERE config_id = $1 ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &values, query, configID)
	return values, err
}

func (r *configValueRepository) FindByID(ctx context.Context, id uuid.UUID) (model.ConfigValues, error) {
	var val model.ConfigValues
	query := `SELECT id, config_id, name, code, created_at, updated_at FROM config_values WHERE id = $1`
	err := r.db.GetContext(ctx, &val, query, id)
	return val, err
}

func (r *configValueRepository) UpsertConfigValue(ctx context.Context, tx *sqlx.Tx, req model.ConfigValues) error {
	query := `
		INSERT INTO config_values (config_id, name, code)
		VALUES ($1, $2, $3)
		ON CONFLICT (code)
		DO UPDATE SET
			name = EXCLUDED.name,
			updated_at = NOW()
	`
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, query, req.ConfigID, req.Name, req.Code)
	} else {
		_, err = r.db.ExecContext(ctx, query, req.ConfigID, req.Name, req.Code)
	}
	return err
}
func (r *configValueRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM config_values WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

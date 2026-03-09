package seeder

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func Run(ctx context.Context, db *sqlx.DB) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if err = seedRoles(ctx, tx, db); err != nil {
		return err
	}
	if err = seedPermissions(ctx, tx, db); err != nil {
		return err
	}
	if err = seedUser(ctx, tx, db); err != nil {
		return err
	}

	return tx.Commit()
}

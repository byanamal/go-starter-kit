package db

import (
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Init(dsn string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}

func WithTransaction(ctx context.Context, db *sqlx.DB, fn func(tx *sqlx.Tx) error) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		slog.ErrorContext(ctx, "db: failed to begin transaction", "error", err)
		return err
	}

	slog.DebugContext(ctx, "db: transaction started")

	defer func() {
		if p := recover(); p != nil {
			slog.WarnContext(ctx, "db: rolling back transaction due to panic", "panic", p)
			if rbErr := tx.Rollback(); rbErr != nil {
				slog.WarnContext(ctx, "db: rollback failed after panic", "error", rbErr)
			}
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		slog.WarnContext(ctx, "db: rolling back transaction", "error", err)
		if rbErr := tx.Rollback(); rbErr != nil {
			slog.WarnContext(ctx, "db: rollback failed", "error", rbErr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		slog.ErrorContext(ctx, "db: failed to commit transaction", "error", err)
		return err
	}

	slog.DebugContext(ctx, "db: transaction committed")
	return nil
}

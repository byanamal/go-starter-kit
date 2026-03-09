-- +migrate Up
ALTER TABLE permissions ADD COLUMN IF NOT EXISTS "group" VARCHAR(100);

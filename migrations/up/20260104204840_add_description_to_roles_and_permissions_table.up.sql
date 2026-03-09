-- +migrate Up
ALTER TABLE roles
ADD COLUMN IF NOT EXISTS description TEXT;

ALTER TABLE permissions
ADD COLUMN IF NOT EXISTS description TEXT;

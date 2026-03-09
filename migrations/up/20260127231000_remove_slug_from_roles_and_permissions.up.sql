-- +migrate Up
ALTER TABLE roles DROP COLUMN IF EXISTS slug;
ALTER TABLE permissions DROP COLUMN IF EXISTS slug;

CREATE UNIQUE INDEX IF NOT EXISTS uq_roles_name ON roles(name);
CREATE UNIQUE INDEX IF NOT EXISTS uq_permissions_name ON permissions(name);
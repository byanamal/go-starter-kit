-- +migrate Down
ALTER TABLE users DROP COLUMN IF EXISTS deleted_at;

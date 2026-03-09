-- +migrate Down
ALTER TABLE permissions DROP COLUMN IF EXISTS "group";

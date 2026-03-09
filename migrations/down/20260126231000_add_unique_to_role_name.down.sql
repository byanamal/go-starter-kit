-- +migrate Down
DROP INDEX IF EXISTS uq_roles_name;

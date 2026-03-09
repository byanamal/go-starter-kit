-- +migrate Down
DROP INDEX IF EXISTS uq_config_code;
DROP INDEX IF EXISTS uq_config_values_code;

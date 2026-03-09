-- +migrate Up
CREATE UNIQUE INDEX IF NOT EXISTS uq_config_code
ON config (code);

CREATE UNIQUE INDEX IF NOT EXISTS uq_config_values_code
ON config_values (code);

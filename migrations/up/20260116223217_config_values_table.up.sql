-- +migrate Up
CREATE TABLE IF NOT EXISTS config_values (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  config_id UUID NOT NULL,
  name VARCHAR(150),
  code TEXT,

  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_cofig
    FOREIGN KEY (config_id)
    REFERENCES config(id)
    ON DELETE CASCADE
);

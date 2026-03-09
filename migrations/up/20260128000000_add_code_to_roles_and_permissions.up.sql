-- Truncate to ensure clean seeding with new code column constraint
TRUNCATE TABLE permissions, roles CASCADE;

ALTER TABLE roles ADD COLUMN IF NOT EXISTS code VARCHAR(100);
ALTER TABLE permissions ADD COLUMN IF NOT EXISTS code VARCHAR(100);

CREATE UNIQUE INDEX IF NOT EXISTS uq_roles_code ON roles(code);
CREATE UNIQUE INDEX IF NOT EXISTS uq_permissions_code ON permissions(code);

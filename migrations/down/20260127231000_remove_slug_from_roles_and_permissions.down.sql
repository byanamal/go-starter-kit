-- +migrate Down
ALTER TABLE roles ADD COLUMN slug VARCHAR(100);
UPDATE roles SET slug = LOWER(REPLACE(name, ' ', '_')) WHERE slug IS NULL;
ALTER TABLE roles ALTER COLUMN slug SET NOT NULL;
ALTER TABLE roles ADD CONSTRAINT roles_slug_key UNIQUE (slug);

ALTER TABLE permissions ADD COLUMN slug VARCHAR(100);
UPDATE permissions SET slug = LOWER(REPLACE(name, ' ', ':')) WHERE slug IS NULL;
ALTER TABLE permissions ALTER COLUMN slug SET NOT NULL;
ALTER TABLE permissions ADD CONSTRAINT permissions_slug_key UNIQUE (slug);

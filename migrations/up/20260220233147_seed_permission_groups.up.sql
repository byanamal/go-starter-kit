-- +migrate Up
UPDATE permissions SET "group" = 'Users' WHERE code LIKE 'user:%';
UPDATE permissions SET "group" = 'Roles' WHERE code LIKE 'role:%';
UPDATE permissions SET "group" = 'Articles' WHERE code LIKE 'article:%';
UPDATE permissions SET "group" = 'Tutorials' WHERE code LIKE 'tutorial:%';
UPDATE permissions SET "group" = 'Dashboards' WHERE code LIKE 'dashboard:%';
UPDATE permissions SET "group" = 'Configs' WHERE code LIKE 'config:%';
UPDATE permissions SET "group" = 'Menus' WHERE code LIKE 'menu:%';
UPDATE permissions SET "group" = 'Permissions' WHERE code LIKE 'permission:%';
UPDATE permissions SET "group" = 'Others' WHERE "group" IS NULL;

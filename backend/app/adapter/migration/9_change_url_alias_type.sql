-- +migrate Up
ALTER TABLE url ALTER COLUMN alias TYPE CHARACTER VARYING(50);

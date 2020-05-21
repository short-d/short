-- +migrate Up
ALTER TABLE "feature_toggle"
    ADD COLUMN "type" VARCHAR(50) DEFAULT 'manual' NOT NULL;

-- +migrate Down
ALTER TABLE "feature_toggle"
    DROP COLUMN "type";
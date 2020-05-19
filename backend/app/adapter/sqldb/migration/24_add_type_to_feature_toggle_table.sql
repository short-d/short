-- +migrate Up
ALTER TABLE "feature_toggle"
    ADD COLUMN "type" VARCHAR(50);

-- +migrate Down
ALTER TABLE "feature_toggle"
    DROP COLUMN "type";
-- +migrate Up
CREATE TABLE "feature_toggle"
(
    "toggle_id"            CHARACTER VARYING(254) NOT NULL PRIMARY KEY,
    "is_enabled"           BIT
);

-- +migrate Down
DROP TABLE "feature_toggle";
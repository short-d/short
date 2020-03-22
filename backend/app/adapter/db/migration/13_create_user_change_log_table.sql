-- +migrate Up
CREATE TABLE "user_change_log"
(
    "user_id"           CHARACTER VARYING(5),
    "last_viewed_at"    TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE "user_change_log";
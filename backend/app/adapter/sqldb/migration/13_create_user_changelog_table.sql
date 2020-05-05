-- +migrate Up
CREATE TABLE "user_changelog"
(
    "user_id"              CHARACTER VARYING(5),
    "email"                CHARACTER VARYING(254) NOT NULL PRIMARY KEY,
    "last_viewed_at"       TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE "user_changelog";

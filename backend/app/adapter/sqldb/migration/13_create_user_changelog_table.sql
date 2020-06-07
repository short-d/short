-- +migrate Up
CREATE TABLE "user_changelog"
(
    "user_id"              CHARACTER VARYING(5),
    -- TODO(issue#824): Add Primary Key Constraint to email on migrate up
    "email"                CHARACTER VARYING(254),
    "last_viewed_at"       TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE "user_changelog";

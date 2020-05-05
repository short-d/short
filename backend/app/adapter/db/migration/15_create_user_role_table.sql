-- +migrate Up
CREATE TABLE "user_role"
(
    "user_id"              CHARACTER VARYING(5) NOT NULL PRIMARY KEY,
    "roles"                BIT VARYING NOT NULL
);

-- +migrate Down
DROP TABLE "user_role";
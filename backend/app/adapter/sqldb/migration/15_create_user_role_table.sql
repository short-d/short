-- +migrate Up
CREATE TABLE "user_role"
(
    "user_id"              CHARACTER VARYING(5) NOT NULL,
    "role"                 CHARACTER VARYING(255) NOT NULL,
    CONSTRAINT "pk_user_role_relation" PRIMARY KEY ("user_id", "role")
);

-- +migrate Down
DROP TABLE "user_role";
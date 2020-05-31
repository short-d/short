-- +migrate Up
ALTER TABLE "user_url_relation"
    DROP CONSTRAINT "user_url_relation_user_email_fkey";
ALTER TABLE "user_url_relation"
    DROP COLUMN user_email;

ALTER TABLE "user"
    DROP CONSTRAINT "User_pkey";
ALTER TABLE "user"
    ADD CONSTRAINT "User_pkey" PRIMARY KEY (id);

ALTER TABLE "user_url_relation"
    ADD CONSTRAINT "user_url_relation_user_id_fkey" FOREIGN KEY (user_id) REFERENCES "user"(id);
ALTER TABLE "user_url_relation"
    ADD CONSTRAINT "pk_user_url_relation" PRIMARY KEY (url_alias, user_id);

-- +migrate Down
ALTER TABLE "user_url_relation"
    DROP CONSTRAINT "pk_user_url_relation";
ALTER TABLE "user_url_relation"
    DROP CONSTRAINT "user_url_relation_user_id_fkey";

ALTER TABLE "user"
    DROP CONSTRAINT "User_pkey";
ALTER TABLE "user"
    ADD CONSTRAINT "User_pkey" PRIMARY KEY (email);

ALTER TABLE "user_url_relation"
    ADD COLUMN user_email CHARACTER VARYING(254);
ALTER TABLE "user_url_relation"
    ADD CONSTRAINT "user_url_relation_user_email_fkey" FOREIGN KEY (user_email) REFERENCES "user"(email);
-- TODO: Add Primary Key Constraint to user_email and url_alias on migrate down
-- TODO: Set user_email to NOT NULL on migrate down

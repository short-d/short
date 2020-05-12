-- +migrate Up
ALTER TABLE "user_role"
    ADD CONSTRAINT "user_id_fkey" FOREIGN KEY (user_id) REFERENCES "user"(id);
ALTER TABLE "user_role"
    DROP CONSTRAINT "pk_user_role_relation";
ALTER TABLE "user_role"
    ADD CONSTRAINT "pk_user_role" PRIMARY KEY (user_id, role);

-- +migrate Down
ALTER TABLE "user_role"
    DROP CONSTRAINT "pk_user_role";
ALTER TABLE "user_role"
    ADD CONSTRAINT "pk_user_role_relation" PRIMARY KEY (user_id, role);
ALTER TABLE "user_role"
    DROP CONSTRAINT "user_id_fkey";
-- +migrate Up
ALTER TABLE "github_sso"
    DROP CONSTRAINT "github_sso_short_user_id_fkey";
ALTER TABLE "github_sso"
    ADD CONSTRAINT "github_sso_short_user_id_fkey"
    FOREIGN KEY (short_user_id) REFERENCES "user"(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE;

-- +migrate Down
ALTER TABLE "github_sso"
    DROP CONSTRAINT "github_sso_short_user_id_fkey";
ALTER TABLE "github_sso"
    ADD CONSTRAINT "github_sso_short_user_id_fkey"
    FOREIGN KEY (short_user_id) REFERENCES "user"(id);
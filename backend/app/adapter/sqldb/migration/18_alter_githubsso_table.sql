-- +migrate Up
ALTER TABLE "github_sso"
    DROP CONSTRAINT "pk_github_oauth_user_relation";
ALTER TABLE "github_sso"
    ADD CONSTRAINT "github_sso_short_user_id_fkey" FOREIGN KEY (short_user_id) REFERENCES "user"(id);

-- +migrate Down
ALTER TABLE "github_sso"
    DROP CONSTRAINT "github_sso_short_user_id_fkey";
ALTER TABLE "github_sso"
    ADD CONSTRAINT "pk_github_oauth_user_relation" PRIMARY KEY (github_user_id, short_user_id);
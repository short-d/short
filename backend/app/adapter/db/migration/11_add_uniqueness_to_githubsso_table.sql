-- +migrate Up
ALTER TABLE github_sso
    ADD CONSTRAINT github_user_id_unique UNIQUE (github_user_id);

ALTER TABLE github_sso
    ADD CONSTRAINT short_user_id_unique UNIQUE (short_user_id);

-- +migrate Down
ALTER TABLE github_sso
    DROP CONSTRAINT github_user_id_unique;

ALTER TABLE github_sso
    drop CONSTRAINT short_user_id_unique;
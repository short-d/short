-- +migrate Up
ALTER TABLE github_oauth_user_relation RENAME TO github_sso;

-- +migrate Down
ALTER TABLE github_sso RENAME TO github_oauth_user_relation;
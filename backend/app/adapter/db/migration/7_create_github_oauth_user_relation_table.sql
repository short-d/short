-- +migrate Up
CREATE TABLE github_oauth_user_relation
(
    github_user_id CHARACTER VARYING(254) NOT NULL,
    short_user_id  CHARACTER VARYING(5)   NOT NULL,
    CONSTRAINT pk_github_oauth_user_relation PRIMARY KEY (github_user_id, short_user_id)
);

-- +migrate Down
DROP TABLE github_oauth_user_relation;
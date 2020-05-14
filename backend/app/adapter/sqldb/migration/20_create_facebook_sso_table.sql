-- +migrate Up
CREATE TABLE facebook_sso
(
    facebook_user_id CHARACTER VARYING(254) NOT NULL UNIQUE,
    short_user_id  CHARACTER VARYING(5) NOT NULL UNIQUE,
    FOREIGN KEY (short_user_id) REFERENCES "user"(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE facebook_sso;
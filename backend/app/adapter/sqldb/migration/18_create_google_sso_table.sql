-- +migrate Up
CREATE TABLE google_sso
(
    google_user_id CHARACTER VARYING(254) NOT NULL UNIQUE,
    short_user_id  CHARACTER VARYING(5)   NOT NULL UNIQUE,
    FOREIGN KEY (short_user_id) REFERENCES "user"(id) ON UPDATE CASCADE
);

-- +migrate Down
DROP TABLE google_sso;
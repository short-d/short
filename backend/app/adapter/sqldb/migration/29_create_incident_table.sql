-- +migrate Up
CREATE TABLE incident
(
    id          CHARACTER VARYING(4)      NOT NULL UNIQUE,
    title       CHARACTER VARYING(100)    NOT NULL UNIQUE,
    created_at  TIMESTAMP WITH TIME ZONE  NOT NULL
);

-- +migrate Down
DROP TABLE incident;

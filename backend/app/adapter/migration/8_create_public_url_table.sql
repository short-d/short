-- +migrate Up
CREATE TABLE public_url (
    alias CHARACTER VARYING(50) NOT NULL,
    FOREIGN KEY (alias) REFERENCES url (alias) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +migrate Down
DROP TABLE public_url;
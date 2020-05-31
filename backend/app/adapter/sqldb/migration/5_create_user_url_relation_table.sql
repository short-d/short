-- +migrate Up
CREATE TABLE user_url_relation
(
    -- TODO: Add Primary Key Constraint to user_email and url_alias on migrate up
    -- TODO: Set user_email to NOT NULL on migrate up
    user_email CHARACTER VARYING(254),
    url_alias  CHARACTER VARYING(50)  NOT NULL,
    FOREIGN KEY (user_email) REFERENCES "user" (email) ON UPDATE CASCADE,
    FOREIGN KEY (url_alias) REFERENCES url (alias) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +migrate Down
DROP TABLE user_url_relation;
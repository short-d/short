-- +migrate Up
CREATE TABLE user_url_relation (
	user_email CHARACTER VARYING(254) NOT NULL,
	url_alias CHARACTER VARYING(50) NOT NULL,
	CONSTRAINT pk_user_url_relation PRIMARY KEY (user_email, url_alias),
	FOREIGN KEY (user_email) REFERENCES "user"(email) ON UPDATE CASCADE,
	FOREIGN KEY (url_alias) REFERENCES url(alias) ON DELETE CASCADE ON UPDATE CASCADE
);
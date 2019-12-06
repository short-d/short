-- +migrate Up
ALTER TABLE "user" ADD id CHARACTER VARYING(5);
ALTER TABLE user_url_relation ADD user_id CHARACTER VARYING(5);

-- +migrate Down
ALTER TABLE user_url_relation DROP user_id;
ALTER TABLE "user" DROP id;
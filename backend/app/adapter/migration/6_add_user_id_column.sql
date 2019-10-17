-- +migrate Up
ALTER TABLE "user" ADD id CHARACTER VARYING(5);
ALTER TABLE user_url_relation ADD user_id CHARACTER VARYING(5);

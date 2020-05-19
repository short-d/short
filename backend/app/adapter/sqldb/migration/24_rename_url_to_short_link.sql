-- +migrate Up
ALTER TABLE public_url RENAME TO public_short_link;
ALTER TABLE url RENAME TO short_link;
ALTER TABLE user_url_relation RENAME TO user_short_link_relation;

ALTER TABLE short_link
    RENAME COLUMN original_url TO original_short_link;

ALTER TABLE user_short_link_relation
    RENAME COLUMN url_alias TO short_link_alias;

-- +migrate Down
ALTER TABLE public_short_link RENAME TO public_url;
ALTER TABLE short_link RENAME TO url;
ALTER TABLE user_short_link_relation RENAME TO user_url_relation;

ALTER TABLE url
    RENAME COLUMN original_short_link TO original_url;

ALTER TABLE user_url_relation
    RENAME COLUMN short_link_alias TO url_alias;
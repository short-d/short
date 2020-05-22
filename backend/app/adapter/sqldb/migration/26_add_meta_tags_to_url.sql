-- +migrate Up
ALTER TABLE "short_link"
    ADD COLUMN "og_title" VARCHAR(200) DEFAULT 'Short: Free link shortening service',
    ADD COLUMN "og_description" VARCHAR(200) DEFAULT 'Short enables people to type less for their favorite web sites',
    ADD COLUMN "og_image_url" VARCHAR(200) DEFAULT 'https://short-d.com/promo/small-tile.png',
    ADD COLUMN "twitter_title" VARCHAR(200) DEFAULT 'Short: Free link shortening service',
    ADD COLUMN "twitter_description" VARCHAR(200) DEFAULT 'Short enables people to type less for their favorite web sites',
    ADD COLUMN "twitter_image_url" VARCHAR(200) DEFAULT 'https://short-d.com/promo/twitter-card.png';

-- +migrate Down
ALTER TABLE "short_link"
    DROP COLUMN "og_title",
    DROP COLUMN "og_description",
    DROP COLUMN "og_image_url",
    DROP COLUMN "twitter_title",
    DROP COLUMN "twitter_description",
    DROP COLUMN "twitter_image_url";

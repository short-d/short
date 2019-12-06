-- +migrate Up
ALTER TABLE "Url" RENAME TO url;
ALTER TABLE url RENAME COLUMN "originalUrl" TO original_url;
ALTER TABLE url ALTER COLUMN original_url TYPE CHARACTER VARYING(200);
ALTER TABLE url RENAME COLUMN "expireAt" TO expire_at;
ALTER TABLE url RENAME COLUMN "createdAt" TO created_at;
ALTER TABLE url RENAME COLUMN "updatedAt" TO updated_at;

-- +migrate Down
ALTER TABLE url RENAME COLUMN updated_at TO "updatedAt";
ALTER TABLE url RENAME COLUMN created_at TO "createdAt";
ALTER TABLE url RENAME COLUMN expire_at TO "expireAt";
ALTER TABLE url ALTER COLUMN original_url TYPE TEXT;
ALTER TABLE url RENAME COLUMN original_url TO "originalUrl";
ALTER TABLE "url" RENAME TO "Url";
-- +migrate Up
ALTER TABLE "User"
RENAME TO "user";

ALTER TABLE "user" RENAME COLUMN "lastSignedInAt" TO last_signed_in_at;
ALTER TABLE "user" RENAME COLUMN "expireAt" TO expire_at;
ALTER TABLE "user" RENAME COLUMN "createdAt" TO created_at;
ALTER TABLE "user" RENAME COLUMN "updatedAt" TO updated_at;
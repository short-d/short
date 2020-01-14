-- +migrate Up
ALTER TABLE "User" RENAME TO "user";
ALTER TABLE "user" RENAME COLUMN "lastSignedInAt" TO last_signed_in_at;
ALTER TABLE "user" RENAME COLUMN "createdAt" TO created_at;
ALTER TABLE "user" RENAME COLUMN "updatedAt" TO updated_at;

-- +migrate Down
ALTER TABLE "user" RENAME COLUMN updated_at TO "updatedAt";
ALTER TABLE "user" RENAME COLUMN created_at TO "createdAt";
ALTER TABLE "user" RENAME COLUMN last_signed_in_at TO "lastSignedInAt";
ALTER TABLE "user" RENAME TO "User";
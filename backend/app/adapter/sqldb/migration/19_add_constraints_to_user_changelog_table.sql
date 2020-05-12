-- +migrate Up
ALTER TABLE "user_changelog"
    ADD CONSTRAINT "user_id_fkey" FOREIGN KEY (user_id) REFERENCES "user"(id);
ALTER TABLE "user_changelog"
    DROP CONSTRAINT "user_changelog_pkey";
ALTER TABLE "user_changelog"
    DROP COLUMN email;
ALTER TABLE "user_changelog"
    ADD CONSTRAINT "user_changelog_pkey" PRIMARY KEY (user_id);

-- +migrate Down
ALTER TABLE "user_changelog"
    DROP CONSTRAINT "user_changelog_pkey";
ALTER TABLE "user_changelog"
    ADD COLUMN email CHARACTER VARYING(254) NOT NULL;
ALTER TABLE "user_changelog"
    ADD CONSTRAINT "user_changelog_pkey" PRIMARY KEY (email);
ALTER TABLE "user_changelog"
    DROP CONSTRAINT "user_id_fkey";
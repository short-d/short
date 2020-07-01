-- +migrate Up
CREATE TABLE "app"
(
    "id" VARCHAR(10) NOT NULL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL
);

-- +migrate Down
DROP TABLE "app";

-- +migrate Up
CREATE TABLE "Url"
(
    "alias"       CHARACTER VARYING PRIMARY KEY,
    "originalUrl" TEXT,
    "expireAt"    TIMESTAMP WITH TIME ZONE,
    "createdAt"   TIMESTAMP WITH TIME ZONE,
    "updatedAt"   TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE "Url";
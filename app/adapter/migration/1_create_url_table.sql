-- +migrate Up
CREATE TABLE "Url" (
    "alias" character varying PRIMARY KEY,
    "originalUrl" text,
    "expireAt" date,
    "createdAt" date,
    "updatedAt" date
);

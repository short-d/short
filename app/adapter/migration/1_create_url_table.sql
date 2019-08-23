-- +migrate Up
CREATE TABLE "Url" (
    "alias" CHARACTER VARYING PRIMARY KEY,
    "originalUrl" TEXT,
    "expireAt" DATA,
    "createdAt" DATA,
    "updatedAt" DATA
);

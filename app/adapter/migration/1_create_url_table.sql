-- +migrate Up
CREATE TABLE "Url" (
    "alias" CHARACTER VARYING PRIMARY KEY,
    "originalUrl" TEXT,
    "expireAt" TIME WITH TIME ZONE,
    "createdAt" TIME WITH TIME ZONE,
    "updatedAt" TIME WITH TIME ZONE
);

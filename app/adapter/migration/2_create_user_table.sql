-- +migrate Up
CREATE TABLE "User" (
    "email" character varying(254) PRIMARY KEY,
    "lastSignInAt" time with time zone,
    "createdAt" time with time zone,
    "expireAt" time with time zone,
    "updatedAt" time with time zone
);

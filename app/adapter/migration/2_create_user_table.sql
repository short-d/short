-- +migrate Up
CREATE TABLE "User" (
    "email" character varying(254) PRIMARY KEY,
    "name" character varying(80),
    "lastSignedInAt" time with time zone,
    "createdAt" time with time zone,
    "expireAt" time with time zone,
    "updatedAt" time with time zone
);

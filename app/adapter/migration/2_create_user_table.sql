-- +migrate Up
CREATE TABLE "User" (
    "email" CHARACTER VARYING(254) PRIMARY KEY,
    "name" CHARACTER VARYING(80),
    "lastSignedInAt" TIME WITH TIME ZONE,
    "createdAt" TIME WITH TIME ZONE,
    "expireAt" TIME WITH TIME ZONE,
    "updatedAt" TIME WITH TIME ZONE
);

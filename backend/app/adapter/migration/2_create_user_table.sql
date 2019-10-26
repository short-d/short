-- +migrate Up
CREATE TABLE "User"
(
    "email"          CHARACTER VARYING(254) PRIMARY KEY,
    "name"           CHARACTER VARYING(80),
    "lastSignedInAt" TIMESTAMP WITH TIME ZONE,
    "createdAt"      TIMESTAMP WITH TIME ZONE,
    "updatedAt"      TIMESTAMP WITH TIME ZONE
);

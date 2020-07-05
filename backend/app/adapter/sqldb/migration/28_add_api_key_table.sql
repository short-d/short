-- +migrate Up
CREATE TABLE "api_key"
(
    "app_id" VARCHAR(10) NOT NULL REFERENCES "app"("id"),
    "key" VARCHAR(10) NOT NULL,
    "disabled" BIT(1) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT pk_api_key PRIMARY KEY ("app_id", "key")
);

-- +migrate Down
DROP TABLE "api_key";

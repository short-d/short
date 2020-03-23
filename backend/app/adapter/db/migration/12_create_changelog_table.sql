-- +migrate Up
CREATE TABLE "change_log"
(
    "id"                CHARACTER VARYING(5) PRIMARY KEY,
    "title"             CHARACTER VARYING (100),
    "summary_markdown"  TEXT,
    "released_at"       TIMESTAMP WITH TIME ZONE
);

-- +migrate Down
DROP TABLE "change_log";
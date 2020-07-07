-- +migrate Up
CREATE TABLE incident (
  id VARCHAR(4) NOT NULL PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL
);
-- +migrate Down
DROP
  TABLE incident;

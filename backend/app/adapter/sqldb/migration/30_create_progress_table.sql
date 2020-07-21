-- +migrate Up
CREATE TABLE progress (
  incident_id VARCHAR(4) NOT NULL,
  FOREIGN KEY (incident_id) REFERENCES "incident"(id) ON DELETE CASCADE ON UPDATE CASCADE,
  status VARCHAR(15) NOT NULL DEFAULT 'reported',
  info TEXT,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL
);
-- +migrate Down
DROP
  TABLE progress;

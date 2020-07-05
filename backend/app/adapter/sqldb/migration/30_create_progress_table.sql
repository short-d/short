-- +migrate Up
CREATE TABLE progress
(
    incident_id CHARACTER VARYING(4) NOT NULL,
    FOREIGN KEY (incident_id) REFERENCES "incident"(id)
        ON UPDATE CASCADE,
    status CHARACTER VARYING(10) NOT NULL DEFAULT 'reported',
    info TEXT,
    created_at TIMESTAMP WITH TIME ZONE  NOT NULL
);

-- +migrate Down
DROP TABLE progress;

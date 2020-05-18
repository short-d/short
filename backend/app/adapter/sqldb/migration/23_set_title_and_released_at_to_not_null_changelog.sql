-- +migrate Up
ALTER TABLE change_log ALTER COLUMN released_at SET NOT NULL;
ALTER TABLE change_log ALTER COLUMN title SET NOT NULL;

-- +migrate Down
ALTER TABLE change_log ALTER COLUMN released_at DROP NOT NULL;
ALTER TABLE change_log ALTER COLUMN title DROP NOT NULL;
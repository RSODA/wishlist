-- +goose Up
ALTER TABLE wish
    ALTER COLUMN url SET DATA TYPE VARCHAR;

-- +goose Down
SELECT 'down SQL query';

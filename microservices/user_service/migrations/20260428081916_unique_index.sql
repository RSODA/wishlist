-- +goose Up
ALTER TABLE sub
    ADD CONSTRAINT uq_id_from_id_to UNIQUE(id_from, id_to);

-- +goose Down

-- +goose Up
CREATE INDEX idx_sub_id_from ON sub(id_from);
CREATE INDEX idx_sub_id_to ON sub(id_to);

-- +goose Down


-- +goose Up
CREATE TABLE sub (
    id_from BIGINT REFERENCES users (tg_id) NOT NULL,
    id_to BIGINT REFERENCES users (tg_id) NOT NULL,
    is_accepted BOOL DEFAULT false
);


-- +goose Down

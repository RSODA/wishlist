-- +goose Up
CREATE TABLE IF NOT EXISTS sub (
                     id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                     id_from BIGINT NOT NULL,
                     id_to   BIGINT NOT NULL,
                     is_accepted BOOLEAN DEFAULT false,
                     FOREIGN KEY (id_from) REFERENCES users(tg_id),
                     FOREIGN KEY (id_to)   REFERENCES users(tg_id)
);

-- +goose Down

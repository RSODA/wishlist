-- +goose Up
CREATE TABLE wish (
                      id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                      tg_id BIGINT NOT NULL,
                      picture VARCHAR(150),
                      title VARCHAR(30),
                      price INT,
                      url VARCHAR(50),
                      status_id INT REFERENCES status(id),
                      status_user_id INT
);

-- +goose Down
DROP TABLE wish;
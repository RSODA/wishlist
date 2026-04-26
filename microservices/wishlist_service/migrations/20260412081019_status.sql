-- +goose Up
CREATE TABLE status (
    id INT GENERATED ALWAYS AS IDENTITY NOT NULL PRIMARY KEY,
    title VARCHAR(15) NOT NULL
);

INSERT INTO status(title) VALUES
('Закажу'),
('Заказал(-а)'),
('Пришло');

-- +goose Down

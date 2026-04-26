-- +goose Up
create table users
(
    tg_id    BIGINT not null UNIQUE PRIMARY KEY,
    username text   not null UNIQUE
);

-- +goose Down
drop table users;
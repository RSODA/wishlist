-- +goose Up
CREATE EXTENSION IF NOT EXISTS pgcrypto;

create table users(
    id        UUID primary key default gen_random_uuid(),
    tg_id     BIGINT not null UNIQUE,
    username  text   not null UNIQUE,
    subscribe JSONB
);

-- +goose Down
drop table users;
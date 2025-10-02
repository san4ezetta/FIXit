-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id         bigserial
    primary key,
    email      varchar(255) not null,
    password   varchar(255) not null,
    name       varchar(255) not null,
    surname    varchar(255) not null,
    patronymic varchar(255)
);

-- +goose Down
DROP TABLE IF EXISTS users;
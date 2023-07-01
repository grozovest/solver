-- +goose Up
CREATE TABLE IF NOT EXISTS users
(
    id                UUID    NOT NULL,
    first_name        VARCHAR NOT NULL,
    second_name       VARCHAR NOT NULL,
    father_name       VARCHAR NOT NULL,
    group_name        VARCHAR NOT NULL,
    password          VARCHAR NOT NULL,
    balance           DECIMAL,
    PRIMARY KEY (id),
    UNIQUE (id)
);

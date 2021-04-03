CREATE DATABASE minera_catalog;

\c minera_catalog;

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(30) NOT NULL,
    password TEXT NOT NULL,
    attempts INT DEFAULT 0
);

CREATE TABLE sessions (
    session_id TEXT NOT NULL
);

CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    added TIMESTAMPTZ NOT NULL
);

CREATE TABLE sub_categories (
    id BIGSERIAL PRIMARY KEY,
    category_id BIGINT NOT NULL REFERENCES categories (id),
    name VARCHAR(50) NOT NULL,
    added TIMESTAMPTZ NOT NULL
);


-- \i C:/Users/kolu/src/minera/catalog/minera_catalog.sql

-- ON DELETE CASCADE

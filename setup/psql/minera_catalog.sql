-- CREATE DATABASE minera_catalog;

-- \c minera_catalog;

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
    added TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE category_paths (
    path VARCHAR(30) PRIMARY KEY,
    function VARCHAR(30) NOT NULL,
    id BIGINT REFERENCES categories (id) NOT NULL
);

INSERT INTO categories (name) VALUES ('Category 1');
INSERT INTO categories (name) VALUES ('Category 2');
INSERT INTO categories (name) VALUES ('Category 3');

INSERT INTO category_paths (path, function, id) VALUES ('/cat/111', 'get_sub_categories', 1);
INSERT INTO category_paths (path, function, id) VALUES ('/cat/112', 'add_category', 1);
INSERT INTO category_paths (path, function, id) VALUES ('/cat/113', 'edit_category', 1);
INSERT INTO category_paths (path, function, id) VALUES ('/cat/114', 'delete_category', 1);

INSERT INTO category_paths (path, function, id) VALUES ('/cat/115', 'get_sub_categories', 2);
INSERT INTO category_paths (path, function, id) VALUES ('/cat/116', 'add_category', 2);
INSERT INTO category_paths (path, function, id) VALUES ('/cat/117', 'edit_category', 2);
INSERT INTO category_paths (path, function, id) VALUES ('/cat/118', 'delete_category', 2);

INSERT INTO category_paths (path, function, id) VALUES ('/cat/119', 'get_sub_categories', 3);
INSERT INTO category_paths (path, function, id) VALUES ('/cat/120', 'add_category', 3);
INSERT INTO category_paths (path, function, id) VALUES ('/cat/121', 'edit_category', 3);
INSERT INTO category_paths (path, function, id) VALUES ('/cat/122', 'delete_category', 3);

-- \i C:/Users/kolu/src/minera/catalog/minera_catalog.sql


-- SELECT name, path, function
-- FROM categories
-- JOIN category_paths
-- ON category_paths.id = categories.id
-- WHERE path = '/cat/111';

-- ON DELETE CASCADE

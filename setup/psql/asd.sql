CREATE DATABASE minera_catalog;

\c minera_catalog;

CREATE TABLE paths (
    path VARCHAR(30) PRIMARY KEY NOT NULL,
    function VARCHAR(30) NOT NULL,
    target_id BIGINT NOT NULL
)

INSERT INTO paths (path, function, target_id) VALUES ('/1', 'list_sub_categories', '1');
INSERT INTO paths (path, function, target_id) VALUES ('/2', 'add_sub_category', '1');
INSERT INTO paths (path, function, target_id) VALUES ('/3', 'rename_categories', '1');
INSERT INTO paths (path, function, target_id) VALUES ('/4', 'delete_categories', '1');

CREATE TABLE categories (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    list_sub_categories VARCHAR(30) NOT NULL,
    add_sub_category VARCHAR(30) NOT NULL,
    rename_category VARCHAR(30) NOT NULL,
    delete_category VARCHAR(30) NOT NULL
);

INSERT INTO categories (name, list_sub_categories, add_sub_category, rename_category, delete_category) VALUES
('Category 1', '/1', '/2', '/3', '/4');

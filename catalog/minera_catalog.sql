CREATE DATABASE minera_catalog;

\c minera_catalog;

CREATE TABLE categories (
    path VARCHAR(30) PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    function VARCHAR(30) NOT NULL,
    added TIMESTAMPTZ
);


CREATE TABLE sub_categories (
    parent VARCHAR(30) REFERENCES categories (path) NOT NULL, 
    path VARCHAR(30) PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    function VARCHAR(30) NOT NULL,
    added TIMESTAMPTZ
);


CREATE TABLE products (
    parent VARCHAR(30) REFERENCES sub_categories (path) NOT NULL, 
    path VARCHAR(30) PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    images TEXT [],
    function VARCHAR(30) NOT NULL,
    added TIMESTAMPTZ
);

-- CREATE TABLE active {
--     path
-- }

-- CREATE TABLE paths (
--     path VARCHAR(50) PRIMARY KEY NOT NULL UNIQUE
-- );


INSERT INTO categories (path, name, function, added)
VALUES ('/cat/asd', 'Category 1', 'get_sub_categories', NOW());

INSERT INTO sub_categories (parent, path, name, function, added)
VALUES ('/cat/asd', '/sub/bsd', 'Sub-Category 1', 'get_products', NOW());

INSERT INTO products (parent, path, name, description, images, function, added)
VALUES ('/sub/bsd', '/pro/csd', 'Product 1', 'Description of Product 1', ARRAY[]::TEXT[], 'delete_product', NOW());

-- ON DELETE CASCADE

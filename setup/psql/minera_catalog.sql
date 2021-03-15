CTEATE DATABASE minera_catalog;

CREATE TABLE categories (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE items (
    category_id BIGINT REFERENCES categories (id) NOT NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    images TEXT [] 
);

CREATE TABLE login (
    username VARCHAR(50) NOT NULL,
    password TEXT NOT NULL,
    attempts INT NOT NULL
);

INSERT INTO categories (name) VALUES ('Category 1');

INSERT INTO items (category_id, name, description, images)
VALUES (1, 'Item 1', 'Description for item 1', ARRAY [ 'images/img1.jpg', 'images/img2.jpg', 'images/img3.jpg']);

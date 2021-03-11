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
    password TEXT NOT NULL
);

INSERT INTO categories (name) VALUES ('Категория 1');

INSERT INTO items (category_id, name, description, images)
VALUES (1, 'Артикул 1', 'Описание на артикул 1', ARRAY [ 'images/img1.jpg', 'images/img2.jpg', 'images/img3.jpg']);

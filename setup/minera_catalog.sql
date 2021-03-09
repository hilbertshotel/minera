-- CTEATE DATABASE minera_catalog;

CREATE TABLE categories (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE items (
    category_id BIGINT REFERENCES categories (id),
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    img1 VARCHAR(100),
    img2 VARCHAR(100),
    img3 VARCHAR(100)
);

CREATE TABLE login (
    password VARCHAR(300) NOT NULL
);

-- INSERT INTO categories (name) VALUES ('Категория 1');
-- INSERT INTO categories (name) VALUES ('Категория 2');


-- INSERT INTO items (category_id, name, description, img1, img2, img3)
-- VALUES (
--     1,
--     'Артикул 1',
--     'Описание на артикул 1',
--     'images/img1.jpg',
--     'images/img2.jpg',
--     'images/img3.jpg'
-- );
-- INSERT INTO items (category_id, name, description, img1, img2)
-- VALUES (
--     1,
--     'Артикул 2',
--     'Описание на артикул 2',
--     'images/img1.jpg',
--     'images/img2.jpg'
-- );
-- INSERT INTO items (category_id, name, description)
-- VALUES (
--     1,
--     'Артикул 3',
--     'Описание на артикул 3'
-- );
-- INSERT INTO items (category_id, name, description, img1, img2, img3)
-- VALUES (
--     2,
--     'Артикул 1',
--     'Описание на артикул 1',
--     'images/img1.jpg',
--     'images/img2.jpg',
--     'images/img3.jpg'
-- );
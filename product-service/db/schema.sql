-- Active: 1659919894211@@127.0.0.1@5432@product@public

DROP TABLE IF EXISTS products;

CREATE TABLE products (
	id serial PRIMARY KEY,
    store_id int NOT NULL,
	name VARCHAR (255) NOT NULL,
	price INT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);
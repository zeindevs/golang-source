# gosqlc

## sqlc generate

```sh
sqlc generate
```

## Create database

```sh
pgsql -U postgres

create database sqlctest;

\c sqlctest

CREATE TABLE IF NOT EXISTS product (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  price NUMERIC(6, 2) NOT NULL,
  available BOOLEAN,
  created timestamp DEFAULT NOW()
);

\dt

INSERT INTO product (name, price, available) VALUES ('Book', 10.99, true);

SELECT * FROM product;

\q
```

-- name: GetProduct :one
SELECT * FROM product WHERE id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM product ORDER BY name;

-- name: CreateProduct :one
INSERT INTO product (
  name, price, available
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM product
WHERE id = $1;

-- name: TotalPrice :one
SELECT SUM(price)::float FROM product;

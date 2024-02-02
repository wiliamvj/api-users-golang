-- name: CreateProduct :exec
INSERT INTO product (id, title, description, price, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6);

-- name: CreateProductCategory :exec
INSERT INTO product_category (id, product_id, category_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5);

-- name: GetCategoryByID :one
SELECT EXISTS (SELECT 1 FROM category WHERE id = $1) AS category_exists;

-- name: GetProductByID :one
SELECT EXISTS (SELECT 1 FROM product WHERE id = $1) AS product_exists;

-- name: UpdateProduct :exec
UPDATE product
SET
title = COALESCE(sqlc.narg('title'), title),
description = COALESCE(sqlc.narg('description'), description),
price = COALESCE(sqlc.narg('price'), price),
updated_at = $2
WHERE id = $1;

-- name: GetCategoriesByProductID :many
SELECT pc.category_id FROM product_category pc WHERE pc.product_id = $1;

-- name: DeleteProductCategory :exec
DELETE FROM product_category WHERE product_id = $1 AND category_id = $2;

-- name: DeleteProduct :exec
DELETE FROM product WHERE id = $1;

-- name: FindManyProducts :many
SELECT
p.id, 
p.title, 
p.description, 
p.price, 
p.created_at
FROM product p 
JOIN product_category pc ON pc.product_id = p.id
WHERE 
  (pc.category_id  = ANY(@categories::TEXT[]) OR @categories::TEXT[] IS NULL)
  AND (
    p.title ILIKE '%' || COALESCE(@search, sqlc.narg('search'), '') || '%' 
    OR 
    p.description ILIKE '%' || COALESCE(@search, sqlc.narg('search'), '') || '%'
  )
ORDER BY p.created_at DESC;

-- name: GetProductCategories :many
SELECT c.id, c.title FROM category c 
JOIN product_category pc ON pc.category_id = c.id 
WHERE pc.product_id = $1;

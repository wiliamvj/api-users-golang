-- name: CreateCategory :exec
INSERT INTO category (id, title, created_at, updated_at) 
VALUES ($1, $2, $3, $4);

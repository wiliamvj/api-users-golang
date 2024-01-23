
-- name: CreateUser :exec
INSERT INTO users (id, name, email, password, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6);

-- name: FindUserByEmail :one
SELECT u.id, u.name, u.email FROM users u WHERE u.email = $1;

-- name: FindUserByID :one
SELECT u.id, u.name, u.email, u.created_At, u.updated_at, a.cep, a.uf, a.city, a.complement, a.street 
FROM users u
JOIN address a ON a.user_id = u.id
WHERE u.id = $1;

-- name: UpdateUser :exec
UPDATE users SET 
name = COALESCE(sqlc.narg('name'), name), 
email = COALESCE(sqlc.narg('email'), email), 
updated_at = $2 
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: FindManyUsers :many
SELECT u.id, u.name, u.email, u.created_At, u.updated_at, a.cep, a.uf, a.city, a.complement, a.street 
FROM users u 
JOIN address a ON a.user_id = u.id
ORDER BY u.created_at DESC;

-- name: UpdatePassword :exec
UPDATE users SET password = $2, updated_at = $3 WHERE id = $1;

-- name: GetUserPassword :one
SELECT u.password FROM users u WHERE u.id = $1;

-- name: CreateUserAddress :exec
INSERT INTO address (id, user_id, cep, ibge, uf, city, complement, street, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: UpdateUserAddress :exec
UPDATE address SET 
cep = COALESCE(sqlc.narg('cep'), cep), 
ibge = COALESCE(sqlc.narg('ibge'), ibge), 
uf = COALESCE(sqlc.narg('uf'), uf), 
city = COALESCE(sqlc.narg('city'), city), 
complement = COALESCE(sqlc.narg('complement'), complement), 
street = COALESCE(sqlc.narg('street'), street), 
updated_at = $2
WHERE user_id = $1;
-- name: GetUserByID :one
SELECT * from users u where u.id = $1;

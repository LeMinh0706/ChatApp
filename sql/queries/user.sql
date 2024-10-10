-- name: CreateUser :one
INSERT INTO users(
    username,
    password,
    url_avatar
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetListUser :many
SELECT * FROM users
WHERE id != $1
LIMIT $2
OFFSET $3;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1
LIMIT 1;
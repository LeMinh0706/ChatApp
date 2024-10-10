-- name: CreateMessage :one
INSERT INTO message (
    from_user,
    to_user,
    content
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetMessages :many
SELECT * FROM message 
WHERE (from_user = $1 AND to_user = $2) OR (to_user = $1 AND from_user = $2)
ORDER BY id DESC;
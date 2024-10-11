-- name: CreateMessage :one
INSERT INTO message (
    from_id,
    to_id,
    content
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetMessages :many
SELECT * FROM message 
WHERE (from_id = $1 AND to_id = $2) OR (to_id = $1 AND from_id = $2)
ORDER BY id DESC;
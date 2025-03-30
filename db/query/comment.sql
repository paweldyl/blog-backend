-- name: CreateComment :one
INSERT INTO comments(
    id,
    text,
    user_id,
    post_id
) VALUES(
    $1, $2, $3, $4
)
RETURNING *;

-- name: GetComment :one
SELECT * FROM comments
WHERE id=$1 LIMIT 1;

-- name: GetCommentForUpdate :one
SELECT * FROM comments
WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: GetPostCommentsWithUsers :many
SELECT
  c.*,
  u.username
FROM comments c
JOIN users u ON c.user_id = u.id
WHERE c.post_id = $1
ORDER BY c.created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateComment :one
UPDATE comments
SET 
    text=$1,
    updated_at = NOW()
WHERE id = $2
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id=$1;
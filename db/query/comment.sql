-- CreateComment :one
INSERT INTO comments(
    id,
    text,
    author_login,
    post_id
) VALUES(
    $1, $2, $3, $4
)
RETURNING *;

-- GetComment :one
SELECT * FROM comments
WHERE id=$1 LIMIT 1;

-- name: GetCommentForUpdate :one
SELECT * FROM comments
WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- GetComments :many
SELECT * FROM comments
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- UpdateComment :one
UPDATE comments
SET 
    text=$1,
    updated_at = NOW()
WHERE id = $2
RETURNING *;

-- DeletePost :exec
DELETE FROM posts
WHERE id=$1;
-- CreatePost :one
INSERT INTO posts(
    id,
    title,
    short_desc,
    description,
    author_login
) VALUES(
    $1, $2, $3, $4, $5
)
RETURNING *;

-- GetPost :one
SELECT * FROM posts
WHERE id=$1 LIMIT 1;

-- name: GetPostForUpdate :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- GetPosts :many
SELECT * FROM posts
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- UpdatePost :one
UPDATE posts
SET 
    title = COALESCE(sqlc.narg('title'), title),
    short_desc = COALESCE(sqlc.narg('short_desc'), short_desc),
    description = COALESCE(sqlc.narg('description'), description),
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- DeletePost :exec
DELETE FROM posts
WHERE id=$1;
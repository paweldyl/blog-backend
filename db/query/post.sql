-- name: CreatePost :one
INSERT INTO posts(
    id,
    title,
    short_desc,
    description,
    user_id
) VALUES(
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetPost :one
SELECT 
  posts.*, 
  users.username 
FROM posts
JOIN users ON posts.user_id = users.id
WHERE posts.id=$1 LIMIT 1;

-- name: GetPostForUpdate :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: GetPostsListing :many
SELECT 
  posts.*, 
  users.username 
FROM posts
JOIN users ON posts.user_id = users.id
ORDER BY posts.created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
SET 
    title = COALESCE(sqlc.narg('title'), title),
    short_desc = COALESCE(sqlc.narg('short_desc'), short_desc),
    description = COALESCE(sqlc.narg('description'), description),
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdateLikesAndDislikes :one
UPDATE posts
SET 
  likes_amount = likes_amount + sqlc.arg(likes_amount),
  dislikes_amount = dislikes_amount + sqlc.arg(dislikes_amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id=$1;
-- name: CreatePostLike :one
INSERT INTO posts_likes(
    user_id,
    post_id,
    value
) VALUES(
    $1, $2, $3
)
RETURNING *;

-- name: GetPostLike :one
SELECT * FROM posts_likes
WHERE user_id=$1 AND post_id=$2 LIMIT 1;

-- name: GetPostLikeForUpdate :one
SELECT * FROM posts_likes
WHERE user_id=$1 AND post_id=$2 LIMIT 1 FOR UPDATE;

-- name: UpdatePostLike :one
UPDATE posts_likes
SET 
    value=$1,
    updated_at = NOW()
WHERE user_id=$2 AND post_id=$3
RETURNING *;

-- name: DeletePostLike :exec
DELETE FROM posts_likes
WHERE user_id=$1 AND post_id=$2;
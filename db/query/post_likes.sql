-- CreatePostLike :one
INSERT INTO posts_likes(
    user_login,
    post_id,
    value
) VALUES(
    $1, $2, $3
)
RETURNING *;

-- GetPostLike :one
SELECT * FROM posts_likes
WHERE user_login=$1 AND post_id=$2 LIMIT 1;

-- name: GetPostLikeForUpdate :one
SELECT * FROM posts_likes
WHERE user_login=$1 AND post_id=$2 LIMIT 1 FOR NO KEY UPDATE;

-- UpdatePostLike :one
UPDATE posts_likes
SET 
    text=$1,
    updated_at = NOW()
WHERE user_login=$1 AND post_id=$2
RETURNING *;

-- DeletePost :exec
DELETE FROM posts
WHERE id=$1;
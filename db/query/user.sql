-- name: CreateUser :one
INSERT INTO users (
  id,
  login,
  hashed_password,
  username
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByLogin :one
SELECT * FROM users
WHERE login = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE login = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: UpdateUser :one
UPDATE users
SET
  username = COALESCE(sqlc.narg('username'), username),
  -- profile_image = COALESCE(sqlc.narg('profile_image'), profile_image),
  updated_at = NOW()
WHERE login = $1
RETURNING *;
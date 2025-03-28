// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  login,
  hashed_password,
  username
) VALUES (
  $1, $2, $3
)
RETURNING login, hashed_password, username, edited_at, created_at
`

type CreateUserParams struct {
	Login          string `json:"login"`
	HashedPassword string `json:"hashed_password"`
	Username       string `json:"username"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Login, arg.HashedPassword, arg.Username)
	var i User
	err := row.Scan(
		&i.Login,
		&i.HashedPassword,
		&i.Username,
		&i.EditedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT login, hashed_password, username, edited_at, created_at FROM users
WHERE login = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, login string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, login)
	var i User
	err := row.Scan(
		&i.Login,
		&i.HashedPassword,
		&i.Username,
		&i.EditedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserForUpdate = `-- name: GetUserForUpdate :one
SELECT login, hashed_password, username, edited_at, created_at FROM users
WHERE login = $1 LIMIT 1 FOR NO KEY UPDATE
`

func (q *Queries) GetUserForUpdate(ctx context.Context, login string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserForUpdate, login)
	var i User
	err := row.Scan(
		&i.Login,
		&i.HashedPassword,
		&i.Username,
		&i.EditedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
  username = COALESCE($2, username),
  -- profile_image = COALESCE(sqlc.narg('profile_image'), profile_image),
  updated_at = NOW()
WHERE login = $1
RETURNING login, hashed_password, username, edited_at, created_at
`

type UpdateUserParams struct {
	Login    string         `json:"login"`
	Username sql.NullString `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.Login, arg.Username)
	var i User
	err := row.Scan(
		&i.Login,
		&i.HashedPassword,
		&i.Username,
		&i.EditedAt,
		&i.CreatedAt,
	)
	return i, err
}

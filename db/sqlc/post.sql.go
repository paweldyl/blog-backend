// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: post.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts(
    id,
    title,
    short_desc,
    description,
    user_id
) VALUES(
    $1, $2, $3, $4, $5
)
RETURNING id, title, short_desc, description, user_id, likes_amount, dislikes_amount, updated_at, created_at
`

type CreatePostParams struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	ShortDesc   string    `json:"short_desc"`
	Description string    `json:"description"`
	UserID      uuid.UUID `json:"user_id"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.Title,
		arg.ShortDesc,
		arg.Description,
		arg.UserID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ShortDesc,
		&i.Description,
		&i.UserID,
		&i.LikesAmount,
		&i.DislikesAmount,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE id=$1
`

func (q *Queries) DeletePost(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const getPost = `-- name: GetPost :one
SELECT id, title, short_desc, description, user_id, likes_amount, dislikes_amount, updated_at, created_at FROM posts
WHERE id=$1 LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ShortDesc,
		&i.Description,
		&i.UserID,
		&i.LikesAmount,
		&i.DislikesAmount,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getPostForUpdate = `-- name: GetPostForUpdate :one
SELECT id, title, short_desc, description, user_id, likes_amount, dislikes_amount, updated_at, created_at FROM posts
WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE
`

func (q *Queries) GetPostForUpdate(ctx context.Context, id uuid.UUID) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPostForUpdate, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ShortDesc,
		&i.Description,
		&i.UserID,
		&i.LikesAmount,
		&i.DislikesAmount,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getPostsListing = `-- name: GetPostsListing :many
SELECT id, title, short_desc, description, user_id, likes_amount, dislikes_amount, updated_at, created_at FROM posts
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
`

type GetPostsListingParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetPostsListing(ctx context.Context, arg GetPostsListingParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPostsListing, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.ShortDesc,
			&i.Description,
			&i.UserID,
			&i.LikesAmount,
			&i.DislikesAmount,
			&i.UpdatedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateLikesAndDislikes = `-- name: UpdateLikesAndDislikes :one
UPDATE posts
SET 
  likes_amount = likes_amount + $1,
  dislikes_amount = dislikes_amount + $2
WHERE id = $3
RETURNING id, title, short_desc, description, user_id, likes_amount, dislikes_amount, updated_at, created_at
`

type UpdateLikesAndDislikesParams struct {
	LikesAmount    int32     `json:"likes_amount"`
	DislikesAmount int32     `json:"dislikes_amount"`
	ID             uuid.UUID `json:"id"`
}

func (q *Queries) UpdateLikesAndDislikes(ctx context.Context, arg UpdateLikesAndDislikesParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updateLikesAndDislikes, arg.LikesAmount, arg.DislikesAmount, arg.ID)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ShortDesc,
		&i.Description,
		&i.UserID,
		&i.LikesAmount,
		&i.DislikesAmount,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET 
    title = COALESCE($2, title),
    short_desc = COALESCE($3, short_desc),
    description = COALESCE($4, description),
  updated_at = NOW()
WHERE id = $1
RETURNING id, title, short_desc, description, user_id, likes_amount, dislikes_amount, updated_at, created_at
`

type UpdatePostParams struct {
	ID          uuid.UUID      `json:"id"`
	Title       sql.NullString `json:"title"`
	ShortDesc   sql.NullString `json:"short_desc"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost,
		arg.ID,
		arg.Title,
		arg.ShortDesc,
		arg.Description,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ShortDesc,
		&i.Description,
		&i.UserID,
		&i.LikesAmount,
		&i.DislikesAmount,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}

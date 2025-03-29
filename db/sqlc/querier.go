// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetCommentForUpdate(ctx context.Context, id uuid.UUID) (Comment, error)
	GetPostForUpdate(ctx context.Context, id uuid.UUID) (Post, error)
	GetPostLikeForUpdate(ctx context.Context, arg GetPostLikeForUpdateParams) (PostsLike, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByLogin(ctx context.Context, login string) (User, error)
	GetUserForUpdate(ctx context.Context, login string) (User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)

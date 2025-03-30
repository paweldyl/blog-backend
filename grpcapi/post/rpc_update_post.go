package post

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/grpcapi/api"
	pb "github.com/paweldyl/blog-backend/pb/post"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	authPayload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, api.UnauthenticatedError(err)
	}

	violations := validateUpdatePostRequest(req)
	if violations != nil {
		return nil, api.InvalidArgumentError(violations)
	}

	postID, err := uuid.Parse(req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid post id: %s", err)
	}
	post, err := server.Store.GetPost(ctx, postID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "post not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to fetch post: %s", err)
	}

	if post.UserID != authPayload.UserID {
		return nil, status.Errorf(codes.Internal, "article doesn't belong to user")
	}

	arg := db.UpdatePostParams{
		ID: postID,
		Title: sql.NullString{
			String: req.GetTitle(),
			Valid:  req.Title != nil,
		},
		ShortDesc: sql.NullString{
			String: req.GetShortDesc(),
			Valid:  req.ShortDesc != nil,
		},
		Description: sql.NullString{
			String: req.GetDescription(),
			Valid:  req.Description != nil,
		},
	}

	updatedPost, err := server.Store.UpdatePost(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "post not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update post: %s", err)
	}

	return &pb.UpdatePostResponse{
		Post: convertPost(updatedPost),
	}, nil
}

func validateUpdatePostRequest(req *pb.UpdatePostRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.Title != nil && len(req.GetTitle()) == 0 {
		err := fmt.Errorf("title can't be empty")
		violations = append(violations, api.FieldViolation("title", err))
	}
	if req.ShortDesc != nil {
		if len(req.GetShortDesc()) < 10 || len(req.GetShortDesc()) > 200 {
			err := fmt.Errorf("short_desc must contain between 10 and 200 signs")
			violations = append(violations, api.FieldViolation("short_desc", err))
		}
	}
	if req.Description != nil && len(req.GetDescription()) < 10 {
		err := fmt.Errorf("description must be at least 10 signs")
		violations = append(violations, api.FieldViolation("description", err))
	}
	return violations
}

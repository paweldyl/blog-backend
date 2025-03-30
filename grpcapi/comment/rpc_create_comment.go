package comment

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/grpcapi/api"
	pb "github.com/paweldyl/blog-backend/pb/comment"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	authPayload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, api.UnauthenticatedError(err)
	}

	violations := validateCreateCommentRequest(req)
	if violations != nil {
		return nil, api.InvalidArgumentError(violations)
	}

	commentID, err := uuid.NewRandom()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate comment ID: %s", err)
	}

	postID, err := uuid.Parse(req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid post id: %s", err)
	}

	//check if post exists
	_, err = server.Store.GetPost(ctx, postID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "post not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to find post: %s", err)
	}

	arg := db.CreateCommentParams{
		ID:     commentID,
		Text:   req.GetText(),
		UserID: authPayload.UserID,
		PostID: postID,
	}

	comment, err := server.Store.CreateComment(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create comment: %s", err)
	}

	return &pb.CreateCommentResponse{
		Comment: convertComment(comment),
	}, nil
}

func validateCreateCommentRequest(req *pb.CreateCommentRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if len(req.GetText()) < 3 {
		err := fmt.Errorf("text must be at least 3 characters")
		violations = append(violations, api.FieldViolation("text", err))
	}
	return violations
}

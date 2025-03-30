package comment

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/grpcapi/api"
	pb "github.com/paweldyl/blog-backend/pb/comment"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateComment(ctx context.Context, req *pb.UpdateCommentRequest) (*pb.UpdateCommentResponse, error) {
	authPayload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, api.UnauthenticatedError(err)
	}

	violations := validateUpdateCommentRequest(req)
	if violations != nil {
		return nil, api.InvalidArgumentError(violations)
	}

	commentID, err := uuid.Parse(req.GetCommentId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid comment id: %s", err)
	}

	comment, err := server.Store.GetComment(ctx, commentID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "comment not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to fetch comment: %s", err)
	}

	if comment.UserID != authPayload.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "comment doesn't belong to user")
	}

	arg := db.UpdateCommentParams{
		ID:   commentID,
		Text: req.GetText(),
	}

	updated, err := server.Store.UpdateComment(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update comment: %s", err)
	}

	return &pb.UpdateCommentResponse{
		Comment: convertComment(updated),
	}, nil
}

func validateUpdateCommentRequest(req *pb.UpdateCommentRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if len(req.GetText()) < 3 {
		err := fmt.Errorf("text must be at least 3 characters")
		violations = append(violations, api.FieldViolation("text", err))
	}
	return violations
}

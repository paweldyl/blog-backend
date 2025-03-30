package comment

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/paweldyl/blog-backend/grpcapi/api"
	pb "github.com/paweldyl/blog-backend/pb/comment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*emptypb.Empty, error) {
	authPayload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, api.UnauthenticatedError(err)
	}

	commentID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid comment id")
	}

	comment, err := server.Store.GetComment(ctx, commentID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "comment not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get comment: %s", err)
	}

	if comment.UserID != authPayload.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "unauthorized to delete this comment")
	}

	err = server.Store.DeleteComment(ctx, commentID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete comment: %s", err)
	}

	return &emptypb.Empty{}, nil
}

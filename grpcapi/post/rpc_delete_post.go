package post

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/paweldyl/blog-backend/grpcapi/api"
	pb "github.com/paweldyl/blog-backend/pb/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*emptypb.Empty, error) {
	authPayload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, api.UnauthenticatedError(err)
	}

	postID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid post id: %s", err)
	}

	post, err := server.Store.GetPost(ctx, postID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "post not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get post: %s", err)
	}

	if post.UserID != authPayload.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "unauthorized to delete this post")
	}

	err = server.Store.DeletePost(ctx, postID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete post: %s", err)
	}

	return &emptypb.Empty{}, nil
}

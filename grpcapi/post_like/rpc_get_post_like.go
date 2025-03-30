package post_like

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/grpcapi/api"
	pb "github.com/paweldyl/blog-backend/pb/post_like"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetPostLike(ctx context.Context, req *pb.GetPostLikeRequest) (*pb.GetPostLikeResponse, error) {
	authPayload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, api.UnauthenticatedError(err)
	}

	postID, err := uuid.Parse(req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid post id: %s", err)
	}
	like, err := server.Store.GetPostLike(ctx, db.GetPostLikeParams{
		UserID: authPayload.UserID,
		PostID: postID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.GetPostLikeResponse{
				PostLike: nil,
			}, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to get post like: %s", err)
	}

	enumVal := mapStringToEnum(string(like.Value))

	return &pb.GetPostLikeResponse{
		PostLike: &pb.PostLike{
			UserId: like.UserID.String(),
			PostId: like.PostID.String(),
			Value:  enumVal,
		},
	}, nil
}

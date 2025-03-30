package post_like

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/grpcapi/api"
	pb "github.com/paweldyl/blog-backend/pb/post_like"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) DeletePostLike(ctx context.Context, req *pb.DeletePostLikeRequest) (*emptypb.Empty, error) {
	authPayload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, api.UnauthenticatedError(err)
	}

	postID, err := uuid.Parse(req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid post id: %s", err)
	}

	err = server.Store.(*db.SQLStore).ExecTx(ctx, func(q *db.Queries) error {
		existing, err := q.GetPostLikeForUpdate(ctx, db.GetPostLikeForUpdateParams{
			UserID: authPayload.UserID,
			PostID: postID,
		})
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil
			}
			return err
		}

		var likesDelta, dislikesDelta int32
		switch existing.Value {
		case "like":
			likesDelta = -1
		case "dislike":
			dislikesDelta = -1
		}

		if _, err := q.UpdateLikesAndDislikes(ctx, db.UpdateLikesAndDislikesParams{
			ID:             postID,
			LikesAmount:    likesDelta,
			DislikesAmount: dislikesDelta,
		}); err != nil {
			return err
		}

		return q.DeletePostLike(ctx, db.DeletePostLikeParams{
			UserID: authPayload.UserID,
			PostID: postID,
		})
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete post like: %s", err)
	}

	return &emptypb.Empty{}, nil
}

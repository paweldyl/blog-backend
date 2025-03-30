// ✅ FILE: grpcapi/post_like/rpc_put_post_like.go
package post_like

import (
	"context"

	"github.com/google/uuid"
	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/grpcapi/api"
	pb "github.com/paweldyl/blog-backend/pb/post_like"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) PutPostLike(ctx context.Context, req *pb.PutPostLikeRequest) (*pb.PutPostLikeResponse, error) {
	authPayload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, api.UnauthenticatedError(err)
	}

	postID, err := uuid.Parse(req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid post id: %s", err)
	}

	newValStr, err := mapEnumToString(req.GetValue())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid like value")
	}

	err = server.Store.(*db.SQLStore).ExecTx(ctx, func(q *db.Queries) error {
		existing, err := q.GetPostLikeForUpdate(ctx, db.GetPostLikeForUpdateParams{
			UserID: authPayload.UserID,
			PostID: postID,
		})

		switch {
		case err == nil:
			// Like exists
			if string(existing.Value) == newValStr {
				return nil
			}

			// Update counters (add new, remove old)
			var likesDelta, dislikesDelta int32
			if existing.Value == "like" {
				likesDelta = -1
				dislikesDelta = 1
			} else {
				likesDelta = 1
				dislikesDelta = -1
			}

			_, err := q.UpdateLikesAndDislikes(ctx, db.UpdateLikesAndDislikesParams{
				ID:             postID,
				LikesAmount:    likesDelta,
				DislikesAmount: dislikesDelta,
			})
			if err != nil {
				return err
			}
			_, err = q.UpdatePostLike(ctx, db.UpdatePostLikeParams{
				UserID: authPayload.UserID,
				PostID: postID,
				Value:  db.LikeValue(newValStr),
			})
			return err

		case err.Error() == "sql: no rows in result set":
			// Like doesn’t exist
			likesDelta := int32(0)
			dislikesDelta := int32(0)
			if newValStr == "like" {
				likesDelta = 1
			} else {
				dislikesDelta = 1
			}

			_, err := q.UpdateLikesAndDislikes(ctx, db.UpdateLikesAndDislikesParams{
				ID:             postID,
				LikesAmount:    likesDelta,
				DislikesAmount: dislikesDelta,
			})
			if err != nil {
				return err
			}
			_, err = q.CreatePostLike(ctx, db.CreatePostLikeParams{
				UserID: authPayload.UserID,
				PostID: postID,
				Value:  db.LikeValue(newValStr),
			})
			return err
		default:
			return err
		}
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to put post like: %s", err)
	}

	return &pb.PutPostLikeResponse{
		PostLike: &pb.PostLike{
			UserId: authPayload.UserID.String(),
			PostId: req.GetPostId(),
			Value:  req.GetValue(),
		},
	}, nil
}

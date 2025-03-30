package post_like

import (
	"fmt"

	db "github.com/paweldyl/blog-backend/db/sqlc"
	pb "github.com/paweldyl/blog-backend/pb/post_like"
)

func convertPostLike(postLike db.PostsLike) *pb.PostLike {
	return &pb.PostLike{
		UserId: postLike.UserID.String(),
		PostId: postLike.PostID.String(),
		Value:  mapStringToEnum(string(postLike.Value)),
	}
}

func mapStringToEnum(value string) pb.LikeValue {
	switch value {
	case "like":
		return pb.LikeValue_LIKE
	case "dislike":
		return pb.LikeValue_DISLIKE
	default:
		return pb.LikeValue_LIKE_VALUE_UNSPECIFIED
	}
}

func mapEnumToString(value pb.LikeValue) (string, error) {
	switch value {
	case pb.LikeValue_LIKE:
		return "like", nil
	case pb.LikeValue_DISLIKE:
		return "dislike", nil
	default:
		return "", fmt.Errorf("invalid like value")
	}
}

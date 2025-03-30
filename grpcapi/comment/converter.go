package comment

import (
	db "github.com/paweldyl/blog-backend/db/sqlc"
	pb "github.com/paweldyl/blog-backend/pb/comment"
)

func convertComment(comment db.Comment) *pb.Comment {
	return &pb.Comment{
		Id:     comment.ID.String(),
		Text:   comment.Text,
		UserId: comment.UserID.String(),
		PostId: comment.PostID.String(),
	}
}

func convertPublicComment(commentWithUser db.GetPostCommentsWithUsersRow) *pb.PublicComment {
	return &pb.PublicComment{
		Id:       commentWithUser.ID.String(),
		Text:     commentWithUser.Text,
		UserId:   commentWithUser.UserID.String(),
		PostId:   commentWithUser.PostID.String(),
		Username: commentWithUser.Username,
	}
}

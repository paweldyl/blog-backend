package post

import (
	db "github.com/paweldyl/blog-backend/db/sqlc"
	pb "github.com/paweldyl/blog-backend/pb/post"
)

func convertPost(post db.Post) *pb.Post {
	return &pb.Post{
		Id:          post.ID.String(),
		Title:       post.Title,
		ShortDesc:   post.ShortDesc,
		Description: post.Description,
		UserId:      post.UserID.String(),
	}
}

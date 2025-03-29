package user

import (
	db "github.com/paweldyl/blog-backend/db/sqlc"
	pb "github.com/paweldyl/blog-backend/pb/user"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username: user.Username,
		Id:       user.ID.String(),
	}
}

package api

import (
	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/pb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username: user.Username,
		Id:       user.ID.String(),
	}
}

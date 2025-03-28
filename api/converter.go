package gapi

import (
	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/pb"
)

func converUser(user db.User) *pb.User {
	return &pb.User{
		Login:    user.Login,
		Username: user.Username,
	}
}

package api

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/lib/pq"
	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/pb"
	"github.com/paweldyl/blog-backend/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	userId, err := uuid.NewRandom()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate uuid: %s", err)
	}
	arg := db.CreateUserParams{
		ID:             userId,
		Login:          req.GetLogin(),
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "user already exist: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "faile to create user: %s", err)
	}

	fmt.Println(user.ID.String())
	fmt.Println(user.Username)
	rsp := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	fmt.Println(rsp)
	return rsp, nil
}

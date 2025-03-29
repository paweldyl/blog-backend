package api

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/paweldyl/blog-backend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}
	userID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id: %s", err)
	}

	user, err := server.store.GetUser(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "error while getting user: %s", err)
	}

	rsp := &pb.GetUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}

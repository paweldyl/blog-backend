package user

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/grpcapi/api"
	pb "github.com/paweldyl/blog-backend/pb/user"
	"github.com/paweldyl/blog-backend/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	authPayload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, api.UnauthenticatedError(err)
	}

	violations := validateUpdateUserRequest(req)
	if violations != nil {
		return nil, api.InvalidArgumentError(violations)
	}

	arg := db.UpdateUserParams{
		ID: authPayload.UserID,
		Username: sql.NullString{
			String: req.GetUsername(),
			Valid:  req.Username != nil,
		},
	}

	user, err := server.Store.UpdateUser(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to update user: %s", err)
	}

	rsp := &pb.UpdateUserResponse{
		User: convertUser(user),
	}
	return rsp, nil
}

func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.Username != nil {
		if err := val.ValidateUsername(req.GetUsername()); err != nil {
			violations = append(violations, api.FieldViolation("username", err))
		}
	}

	return violations
}

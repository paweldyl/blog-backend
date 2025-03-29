package post

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/grpcapi/api"
	pb "github.com/paweldyl/blog-backend/pb/post"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	authPayload, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, api.UnauthenticatedError(err)
	}

	violations := validateCreatePostRequest(req)
	if violations != nil {
		return nil, api.InvalidArgumentError(violations)
	}

	postId, err := uuid.NewRandom()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate uuid: %s", err)
	}
	arg := db.CreatePostParams{
		ID:          postId,
		Title:       req.GetTitle(),
		ShortDesc:   req.GetShortDesc(),
		Description: req.GetDescription(),
		UserID:      authPayload.UserID,
	}

	post, err := server.Store.CreatePost(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "faile to create user: %s", err)
	}

	rsp := &pb.CreatePostResponse{
		Post: convertPost(post),
	}
	fmt.Println(rsp)
	return rsp, nil
}

func validateCreatePostRequest(req *pb.CreatePostRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if len(req.GetTitle()) == 0 {
		err := fmt.Errorf("title can't be empty")
		violations = append(violations, api.FieldViolation("title", err))
	}
	if len(req.GetShortDesc()) < 10 || len(req.GetShortDesc()) > 200 {
		err := fmt.Errorf("short_desc must contain between 10 and 200 signs")
		violations = append(violations, api.FieldViolation("short_desc", err))
	}
	if len(req.GetDescription()) < 10 {
		err := fmt.Errorf("description must be at least 10 signs")
		violations = append(violations, api.FieldViolation("description", err))
	}

	return violations
}

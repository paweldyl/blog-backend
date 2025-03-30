package post

import (
	"context"
	"fmt"

	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/grpcapi/api"
	pb "github.com/paweldyl/blog-backend/pb/post"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetPostsListing(ctx context.Context, req *pb.GetPostsListingRequest) (*pb.GetPostsListingResponse, error) {
	_, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, api.UnauthenticatedError(err)
	}

	violations := validateGetPostsListingRequest(req)
	if violations != nil {
		return nil, api.InvalidArgumentError(violations)
	}

	arg := db.GetPostsListingParams{
		Limit:  req.GetPerPage(),
		Offset: req.GetPage() * req.GetPerPage(),
	}

	posts, err := server.Store.GetPostsListing(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list posts: %s", err)
	}

	var pbPosts []*pb.Post
	for _, post := range posts {
		pbPosts = append(pbPosts, convertPost(post))
	}

	return &pb.GetPostsListingResponse{
		Posts: pbPosts,
	}, nil
}

func validateGetPostsListingRequest(req *pb.GetPostsListingRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.GetPerPage() < 1 {
		err := fmt.Errorf("posts per page must be greater than zero")
		violations = append(violations, api.FieldViolation("per_page", err))
	}
	if req.GetPage() < 0 {
		err := fmt.Errorf("page can't be negative")
		violations = append(violations, api.FieldViolation("page", err))
	}

	return violations
}

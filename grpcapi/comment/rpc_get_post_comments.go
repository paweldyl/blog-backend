package comment

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/grpcapi/api"
	pb "github.com/paweldyl/blog-backend/pb/comment"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetPostComments(ctx context.Context, req *pb.GetPostCommentsRequest) (*pb.GetPostCommentsResponse, error) {
	_, err := server.AuthorizeUser(ctx)
	if err != nil {
		return nil, api.UnauthenticatedError(err)
	}

	violations := validateGetPostCommentsRequest(req)
	if violations != nil {
		return nil, api.InvalidArgumentError(violations)
	}

	postID, err := uuid.Parse(req.GetPostId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid post id: %s", err)
	}

	limit := int(req.GetPerPage())
	// added to check if next page will exist.
	realLimit := limit + 1

	arg := db.GetPostCommentsWithUsersParams{
		PostID: postID,
		Limit:  int32(realLimit),
		Offset: req.GetPage() * req.GetPerPage(),
	}

	comments, err := server.Store.GetPostCommentsWithUsers(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get comments: %s", err)
	}

	nextPageExists := len(comments) > limit
	if nextPageExists {
		comments = comments[:limit]
	}

	var pbComments []*pb.PublicComment
	for _, c := range comments {
		pbComments = append(pbComments, convertPublicComment(c))
	}

	return &pb.GetPostCommentsResponse{
		Comments:       pbComments,
		NextPageExists: nextPageExists,
	}, nil
}

func validateGetPostCommentsRequest(req *pb.GetPostCommentsRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.GetPerPage() < 1 {
		err := fmt.Errorf("per_page must be greater than zero")
		violations = append(violations, api.FieldViolation("per_page", err))
	}
	if req.GetPage() < 0 {
		err := fmt.Errorf("page can't be negative")
		violations = append(violations, api.FieldViolation("page", err))
	}
	return violations
}

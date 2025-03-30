package comment

import (
	"fmt"

	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/grpcapi/api"
	pb "github.com/paweldyl/blog-backend/pb/comment"
	"github.com/paweldyl/blog-backend/token"
	"github.com/paweldyl/blog-backend/util"
)

// Server service gRPC requests for our blog service
type Server struct {
	pb.UnimplementedCommentServiceServer
	api.BaseServer
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		BaseServer: api.BaseServer{
			Config:     config,
			Store:      store,
			TokenMaker: tokenMaker,
		},
	}

	return server, nil
}

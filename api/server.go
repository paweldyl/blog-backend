package api

import (
	"fmt"

	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/pb"
	"github.com/paweldyl/blog-backend/token"
	"github.com/paweldyl/blog-backend/util"
)

// Server service gRPC requests for our blog service
type Server struct {
	// db is the database connection
	pb.UnimplementedBlogServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

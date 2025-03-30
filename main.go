package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/grpcapi/comment"
	"github.com/paweldyl/blog-backend/grpcapi/post"
	"github.com/paweldyl/blog-backend/grpcapi/user"
	commentpb "github.com/paweldyl/blog-backend/pb/comment"
	postpb "github.com/paweldyl/blog-backend/pb/post"
	userpb "github.com/paweldyl/blog-backend/pb/user"
	"github.com/paweldyl/blog-backend/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	userServer, err := user.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	postServer, err := post.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	commentServer, err := comment.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, userServer)
	postpb.RegisterPostServiceServer(grpcServer, postServer)
	commentpb.RegisterCommentServiceServer(grpcServer, commentServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

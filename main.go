package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"

	_ "github.com/golang/mock/gomock"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	_ "github.com/lib/pq"
	db "github.com/paweldyl/blog-backend/db/sqlc"
	"github.com/paweldyl/blog-backend/grpcapi/comment"
	"github.com/paweldyl/blog-backend/grpcapi/post"
	"github.com/paweldyl/blog-backend/grpcapi/post_like"
	"github.com/paweldyl/blog-backend/grpcapi/user"
	commentpb "github.com/paweldyl/blog-backend/pb/comment"
	postpb "github.com/paweldyl/blog-backend/pb/post"
	postlikepb "github.com/paweldyl/blog-backend/pb/post_like"
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
	postLikeServer, err := post_like.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, userServer)
	postpb.RegisterPostServiceServer(grpcServer, postServer)
	commentpb.RegisterCommentServiceServer(grpcServer, commentServer)
	postlikepb.RegisterPostLikeServiceServer(grpcServer, postLikeServer)
	reflection.Register(grpcServer)

	// Start native gRPC server for postman etc.
	go func() {
		listener, err := net.Listen("tcp", config.ServerAddress)
		if err != nil {
			log.Fatalf("failed to listen on %v: %v", config.ServerAddress, err)
		}
		log.Printf("gRPC server running at %v (native gRPC)", config.ServerAddress)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// Start gRPC-Web server for frontend
	wrappedGrpc := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(origin string) bool {
		return true
	}))

	httpServer := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if wrappedGrpc.IsGrpcWebRequest(r) || wrappedGrpc.IsAcceptableGrpcCorsRequest(r) {
			wrappedGrpc.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	log.Printf("gRPC-Web server running at %v", config.WebServerAddress)
	if err := http.ListenAndServe(config.WebServerAddress, httpServer); err != nil {
		log.Fatalf("failed to start gRPC-Web HTTP server: %v", err)
	}
}

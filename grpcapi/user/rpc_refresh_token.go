package user

import (
	"context"
	"errors"
	"time"

	db "github.com/paweldyl/blog-backend/db/sqlc"
	pb "github.com/paweldyl/blog-backend/pb/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	refreshPayload, err := server.TokenMaker.VerifyToken(req.GetRefreshToken())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid refresh token: %v", err)
	}

	session, err := server.Store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "session not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get session: %v", err)
	}

	if session.IsBlocked {
		return nil, status.Errorf(codes.PermissionDenied, "session is blocked")
	}

	if session.UserID != refreshPayload.UserID {
		return nil, status.Errorf(codes.PermissionDenied, "incorrect session user")
	}

	if session.RefreshToken != req.GetRefreshToken() {
		return nil, status.Errorf(codes.PermissionDenied, "mismatched refresh token")
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, status.Errorf(codes.PermissionDenied, "refresh token expired")
	}

	accessToken, accessPayload, err := server.TokenMaker.CreateToken(
		refreshPayload.UserID,
		server.Config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token: %v", err)
	}

	return &pb.RefreshTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.ExpiredAt),
	}, nil
}

package user

import (
	"context"
	"database/sql"

	db "github.com/paweldyl/blog-backend/db/sqlc"
	pb "github.com/paweldyl/blog-backend/pb/user"
	"github.com/paweldyl/blog-backend/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.Store.GetUserByLogin(ctx, req.GetLogin())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "error while getting user: %s", err)
	}

	err = util.CheckPassword(req.GetPassword(), user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "password doesnt match: %s", err)
	}

	accessToken, accessPayload, err := server.TokenMaker.CreateToken(
		user.ID,
		server.Config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while generating access token: %s", err)
	}

	refreshToken, refreshPayload, err := server.TokenMaker.CreateToken(
		user.ID,
		server.Config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while generating refresh token: %s", err)
	}

	session, err := server.Store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserID:       user.ID,
		RefreshToken: refreshToken,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while creating session: %s", err)
	}

	rsp := &pb.LoginUserResponse{
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
		User:                  convertUser(user),
	}
	return rsp, nil
}

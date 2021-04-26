package service

import (
	"context"

	"github.jiaget.com/pc-book/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer is the server for authentication
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	userStore  UserStore
	jwtManager *JWTManager
}

// NewAuthServer returns a new auth server
func NewAuthServer(userStore UserStore, jwtManager *JWTManager) *AuthServer {
	authServer := &AuthServer{
		userStore:  userStore,
		jwtManager: jwtManager,
	}
	return authServer
}

// Login is a unary RPC
func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := server.userStore.Find(req.GetUsername())
	if err != nil {
		return nil, err
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, err
	}

	res := &pb.LoginResponse{
		AccessToken: token,
	}
	return res, nil
}

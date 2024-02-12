package adapters

import (
	"context"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/user/entities"
	"github.com/Chystik/pass-man/internal/user/usecases"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userHandlers struct {
	usecases usecases.UserUsecases
	pb.UnimplementedUserServiceServer
	jwtKey []byte
}

func NewUserHandlers(uu usecases.UserUsecases, jwtKey []byte) *userHandlers {
	return &userHandlers{
		usecases: uu,
		jwtKey:   jwtKey,
	}
}

func (uh *userHandlers) SignUp(ctx context.Context, u *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	var response pb.SignUpResponse

	if err := uh.usecases.CreateUser(ctx, u.User.Login, u.User.Password); err != nil {
		return nil, status.Errorf(codes.Internal, "create user error: %s", err.Error())
	}

	token, err := uh.authorize(u.User.Login)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authorize user error: %s", err.Error())
	}

	response.Token = &pb.Token{Token: token}

	return &response, nil
}

func (uh *userHandlers) Login(ctx context.Context, u *pb.LoginRequest) (*pb.LoginResponse, error) {
	var response pb.LoginResponse

	err := uh.usecases.AuthenticateUser(ctx, u.User.Login, u.User.Password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authenticate user error: %s", err.Error())
	}

	token, err := uh.authorize(u.User.Login)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authorize user error: %s", err.Error())
	}

	response.Token = &pb.Token{Token: token}

	return &response, nil
}

func (uh *userHandlers) authorize(login string) (string, error) {
	claims := entities.AuthClaims{
		Login: login,
	}

	token, err := claims.AuthorizeUser(uh.jwtKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

package adapters

import (
	"context"
	"errors"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/user/entities"

	"google.golang.org/grpc"
)

type UserAPIClient interface {
	SignUp(ctx context.Context, login string, password []byte) (entities.JWTtoken, error)
	Login(ctx context.Context, login string, password []byte) (entities.JWTtoken, error)
}

type userAPIClient struct {
	conn   *grpc.ClientConn
	client pb.UserServiceClient
}

func NewUserAPIClient(conn *grpc.ClientConn, client pb.UserServiceClient) *userAPIClient {
	return &userAPIClient{
		conn:   conn,
		client: client,
	}
}

func (uc *userAPIClient) SignUp(ctx context.Context, login string, password []byte) (entities.JWTtoken, error) {
	user := &pb.User{
		Login:    login,
		Password: password,
	}

	req := &pb.SignUpRequest{
		User: user,
	}

	res, err := uc.client.SignUp(ctx, req)
	if err != nil {
		return "", err
	}

	if res.Error != nil {
		return "", errors.New(res.Error.String())
	}

	return entities.JWTtoken(res.Token.Token), nil
}

func (uc *userAPIClient) Login(ctx context.Context, login string, password []byte) (entities.JWTtoken, error) {
	user := &pb.User{
		Login:    login,
		Password: password,
	}

	req := &pb.LoginRequest{
		User: user,
	}

	res, err := uc.client.Login(ctx, req)
	if err != nil {
		return "", err
	}

	if res.Error != nil {
		return "", errors.New(res.Error.String())
	}

	return entities.JWTtoken(res.Token.Token), nil
}

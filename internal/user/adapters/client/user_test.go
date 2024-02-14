package adapters

import (
	"context"
	"errors"
	"testing"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	mocks "github.com/Chystik/pass-man/internal/infrastructure/grpc/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type userAPIClientMock struct {
	c *mocks.UserServiceClient
}

func TestSignUp_ReturnsNoErr(t *testing.T) {
	t.Parallel()

	mk, cl := newUseAPIClientMock()
	resp := &pb.SignUpResponse{
		Token: &pb.Token{},
	}
	ctx := context.Background()

	mk.c.EXPECT().SignUp(mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
	_, err := cl.SignUp(ctx, "login", []byte("password"))

	assert.NoError(t, err)
}

func TestSignUp_WhenClientSignUpReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newUseAPIClientMock()
	ctx := context.Background()

	mk.c.EXPECT().SignUp(mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("err"))
	_, err := cl.SignUp(ctx, "login", []byte("password"))

	assert.Error(t, err)
}

func TestSignUp_WhenRespReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newUseAPIClientMock()
	resp := &pb.SignUpResponse{
		Error: &pb.Error{},
	}
	ctx := context.Background()

	mk.c.EXPECT().SignUp(mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
	_, err := cl.SignUp(ctx, "login", []byte("password"))

	assert.Error(t, err)
}

func TestLogin_ReturnsNoErr(t *testing.T) {
	t.Parallel()

	mk, cl := newUseAPIClientMock()
	resp := &pb.LoginResponse{
		Token: &pb.Token{},
	}
	ctx := context.Background()

	mk.c.EXPECT().Login(mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
	_, err := cl.Login(ctx, "login", []byte("password"))

	assert.NoError(t, err)
}

func TestLogin_WhenClientSignUpReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newUseAPIClientMock()
	ctx := context.Background()

	mk.c.EXPECT().Login(mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("err"))
	_, err := cl.Login(ctx, "login", []byte("password"))

	assert.Error(t, err)
}

func TestLogin_WhenRespReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newUseAPIClientMock()
	resp := &pb.LoginResponse{
		Error: &pb.Error{},
	}
	ctx := context.Background()

	mk.c.EXPECT().Login(mock.Anything, mock.Anything, mock.Anything).Return(resp, nil)
	_, err := cl.Login(ctx, "login", []byte("password"))

	assert.Error(t, err)
}

func newUseAPIClientMock() (*userAPIClientMock, *userAPIClient) {
	m := &userAPIClientMock{
		c: &mocks.UserServiceClient{},
	}
	return m, NewUserAPIClient(m.c)
}

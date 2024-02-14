package adapters

import (
	"context"
	"errors"
	"testing"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type userHandlersMock struct {
	uc *mocks.UserUsecases
	k  []byte
}

func TestSignUp_ReturnsResult(t *testing.T) {
	t.Parallel()

	mks, uh := newUserHandlersMks()
	ctx := context.Background()
	req := &pb.SignUpRequest{
		User: getTestPbUser(),
	}

	mks.uc.EXPECT().CreateUser(mock.Anything, mock.Anything, mock.Anything).Return(nil)
	_, err := uh.SignUp(ctx, req)

	assert.NoError(t, err)
}

func TestSignUp_WhenCreateUserReturnsErr(t *testing.T) {
	t.Parallel()

	mks, uh := newUserHandlersMks()
	ctx := context.Background()
	req := &pb.SignUpRequest{
		User: getTestPbUser(),
	}

	mks.uc.EXPECT().CreateUser(mock.Anything, mock.Anything, mock.Anything).Return(errors.New("err"))
	_, err := uh.SignUp(ctx, req)

	assert.Error(t, err)
}

func TestLogin_ReturnsResult(t *testing.T) {
	t.Parallel()

	mks, uh := newUserHandlersMks()
	ctx := context.Background()
	req := &pb.LoginRequest{
		User: getTestPbUser(),
	}

	mks.uc.EXPECT().AuthenticateUser(mock.Anything, mock.Anything, mock.Anything).Return(nil)
	_, err := uh.Login(ctx, req)

	assert.NoError(t, err)
}

func TestLogin_WhenAuthenticateReturnsErr(t *testing.T) {
	t.Parallel()

	mks, uh := newUserHandlersMks()
	ctx := context.Background()
	req := &pb.LoginRequest{
		User: getTestPbUser(),
	}

	mks.uc.EXPECT().AuthenticateUser(mock.Anything, mock.Anything, mock.Anything).Return(errors.New("err"))
	_, err := uh.Login(ctx, req)

	assert.Error(t, err)
}

func newUserHandlersMks() (*userHandlersMock, *userHandlers) {
	key := []byte("test")
	m := &userHandlersMock{
		uc: &mocks.UserUsecases{},
		k:  key,
	}

	return m, NewUserHandlers(m.uc, key)
}

func getTestPbUser() *pb.User {
	return &pb.User{
		Login:    "test",
		Password: []byte("test"),
	}
}

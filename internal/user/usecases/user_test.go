package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/Chystik/pass-man/internal/user/entities"
	"github.com/Chystik/pass-man/internal/user/mocks"
	vaultmk "github.com/Chystik/pass-man/internal/vault/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser_WhenRepoReturnsErr(t *testing.T) {
	t.Parallel()
	usecases, mks := getUserUsecasesMocks()

	user := getTestUser()
	pass := []byte{0, 1}
	ctx := context.Background()

	mks.repo.EXPECT().Create(ctx, mock.Anything).Return(errors.New("err"))

	err := usecases.CreateUser(ctx, user.Login, pass)
	assert.Error(t, err)
}

func TestCreateUser_WhenUserVaultUnlockReturnsErr(t *testing.T) {
	t.Parallel()
	usecases, mks := getUserUsecasesMocks()

	user := getTestUser()
	pass := []byte{0, 1}
	ctx := context.Background()

	mks.repo.EXPECT().Create(ctx, mock.Anything).Return(nil)
	mks.ks.EXPECT().Unlock(user.Login, mock.Anything).Return(errors.New("err"))

	err := usecases.CreateUser(ctx, user.Login, pass)
	assert.Error(t, err)
}

func TestCreateUser_ReturnsNoErr(t *testing.T) {
	t.Parallel()
	usecases, mks := getUserUsecasesMocks()

	user := getTestUser()
	pass := []byte{0, 1}
	ctx := context.Background()

	mks.repo.EXPECT().Create(ctx, mock.Anything).Return(nil)
	mks.ks.EXPECT().Unlock(user.Login, mock.Anything).Return(nil)

	err := usecases.CreateUser(ctx, user.Login, pass)
	assert.NoError(t, err)
}

func TestAuthenticateUser_WhenRepoReturnsErr(t *testing.T) {
	t.Parallel()
	usecases, mks := getUserUsecasesMocks()

	user := getTestUser()
	pass := []byte{0, 1}
	ctx := context.Background()

	mks.repo.EXPECT().Get(ctx, mock.Anything).Return(entities.User{}, errors.New("err"))

	err := usecases.AuthenticateUser(ctx, user.Login, pass)
	assert.Error(t, err)
}

func TestAuthenticateUser_WhenVaultUnlockReturnsErr(t *testing.T) {
	t.Parallel()
	usecases, mks := getUserUsecasesMocks()

	user := getTestUser()
	pass := "super pass"
	ctx := context.Background()

	mks.repo.EXPECT().Get(ctx, mock.Anything).Return(user, nil)
	mks.ks.EXPECT().Unlock(user.Login, mock.Anything).Return(errors.New("err"))

	err := usecases.AuthenticateUser(ctx, user.Login, []byte(pass))
	assert.Error(t, err)
}

func TestAuthenticateUser_ReturnsNoErr(t *testing.T) {
	t.Parallel()
	usecases, mks := getUserUsecasesMocks()

	user := getTestUser()
	pass := []byte("super pass")
	ctx := context.Background()

	user.SetPassword(pass)
	user.SetVaultKey(pass)

	mks.repo.EXPECT().Get(ctx, mock.Anything).Return(user, nil)
	mks.ks.EXPECT().Unlock(user.Login, mock.Anything).Return(nil)

	err := usecases.AuthenticateUser(ctx, user.Login, pass)
	assert.NoError(t, err)
}

type userUsecasesMocks struct {
	repo *mocks.UserRepository
	ks   *vaultmk.VaultKeyStore
}

func getUserUsecasesMocks() (*userUsecases, *userUsecasesMocks) {
	m := &userUsecasesMocks{
		repo: &mocks.UserRepository{},
		ks:   &vaultmk.VaultKeyStore{},
	}
	u := NewUserUsecases(m.repo, m.ks)

	return u, m
}

func getTestUser() entities.User {
	return entities.User{
		Login: "test",
	}
}

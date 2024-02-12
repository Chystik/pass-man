package usecases

import (
	"context"

	"github.com/Chystik/pass-man/internal/user/entities"
	"github.com/Chystik/pass-man/internal/vault"
)

type userUsecases struct {
	repo  UserRepository
	vault vault.VaultKeyStore
}

func NewUserUsecases(ur UserRepository, ks vault.VaultKeyStore) *userUsecases {
	return &userUsecases{
		repo:  ur,
		vault: ks,
	}
}

func (u *userUsecases) CreateUser(ctx context.Context, login string, password []byte) error {
	var user entities.User

	user.Login = login

	if err := user.SetPassword(password); err != nil {
		return err
	}

	if err := user.SetVaultKey(password); err != nil {
		return err
	}

	if err := u.repo.Create(ctx, user); err != nil {
		return err
	}

	return u.unlockUserVault(user, password)
}

func (u *userUsecases) AuthenticateUser(ctx context.Context, login string, password []byte) error {
	actual, err := u.repo.Get(ctx, login)
	if err != nil {
		return err
	}

	err = actual.Authenticate(password)
	if err != nil {
		return err
	}

	return u.unlockUserVault(actual, password)
}

func (u *userUsecases) Update(ctx context.Context, user entities.User) error {
	panic("not implemented") // TODO: Implement
}

func (u *userUsecases) Delete(ctx context.Context, user entities.User) error {
	panic("not implemented") // TODO: Implement
}

func (u *userUsecases) unlockUserVault(user entities.User, password []byte) error {
	key, err := user.GetVaultKey(password)
	if err != nil {
		return err
	}

	return u.vault.Unlock(user.Login, key)
}

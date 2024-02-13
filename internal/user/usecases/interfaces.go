package usecases

import (
	"context"

	"github.com/Chystik/pass-man/internal/user/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user entities.User) error
	Get(ctx context.Context, login string) (entities.User, error)
}

type User interface {
	SetPassword(password []byte) error
	Authenticate(password []byte) error
	SetVaultKey(vaultPassword []byte) error
	GetVaultKey(vaultPassword []byte) ([]byte, error)
}

type AuthClaims interface {
	AuthorizeUser(jwtKey []byte) (string, error)
}

package usecases

import (
	"context"

	"github.com/Chystik/pass-man/internal/user/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user entities.User) error
	Get(ctx context.Context, login string) (entities.User, error)
	Update(ctx context.Context, user entities.User) error
	Delete(ctx context.Context, user entities.User) error
}

type UserUsecases interface {
	CreateUser(ctx context.Context, login string, password []byte) error
	AuthenticateUser(ctx context.Context, login string, password []byte) error
}

type UserAPIClient interface {
	SignUp(ctx context.Context, login string, password []byte) (entities.JWTtoken, error)
	Login(ctx context.Context, login string, password []byte) (entities.JWTtoken, error)
}

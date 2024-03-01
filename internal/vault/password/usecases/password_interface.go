package usecases

import (
	"context"

	"github.com/Chystik/pass-man/internal/vault/password/entities"
)

type PasswordRepository interface {
	Create(ctx context.Context, userID string, password entities.Password) error
	GetOne(ctx context.Context, userID string, meta string) (entities.Password, error)
	Delete(ctx context.Context, userID string, meta string) error
	GetList(ctx context.Context, userID string) ([]entities.Password, error)
}

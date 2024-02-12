package usecases

import (
	"context"

	"github.com/Chystik/pass-man/internal/vault/entities"
)

type passwordUsecases struct {
	passwordRepo PasswordRepository
}

func NewPasswordUsecases(p PasswordRepository) *passwordUsecases {
	return &passwordUsecases{
		passwordRepo: p,
	}
}

func (vu *passwordUsecases) AddPassword(ctx context.Context, userID string, password entities.Password) error {
	return vu.passwordRepo.Create(ctx, userID, password)
}

func (vu *passwordUsecases) GetPassword(ctx context.Context, userID string, meta string) (entities.Password, error) {
	return vu.passwordRepo.GetOne(ctx, userID, meta)
}

func (vu *passwordUsecases) DeletePassword(ctx context.Context, userID string, meta string) error {
	return vu.passwordRepo.Delete(ctx, userID, meta)
}

func (vu *passwordUsecases) ListPassword(ctx context.Context, userID string) ([]entities.Password, error) {
	return vu.passwordRepo.GetList(ctx, userID)
}

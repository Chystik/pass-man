package usecases

import (
	"context"

	"github.com/Chystik/pass-man/internal/vault/card/entities"
)

type CardRepository interface {
	Create(ctx context.Context, userID string, card entities.Card) error
	GetOne(ctx context.Context, userID string, meta string) (entities.Card, error)
	Delete(ctx context.Context, userID string, meta string) error
	GetList(ctx context.Context, userID string) ([]entities.Card, error)
}

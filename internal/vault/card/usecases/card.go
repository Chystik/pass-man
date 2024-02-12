package usecases

import (
	"context"

	"github.com/Chystik/pass-man/internal/vault/card/entities"
)

type CardUsecases interface {
	AddCard(ctx context.Context, userID string, card entities.Card) error
	GetCard(ctx context.Context, userID string, meta string) (entities.Card, error)
	DeleteCard(ctx context.Context, userID string, meta string) error
	ListCard(ctx context.Context, userID string) ([]entities.Card, error)
}

type cardUsecases struct {
	cardRepo CardRepository
}

func NewCardUsecases(c CardRepository) *cardUsecases {
	return &cardUsecases{
		cardRepo: c,
	}
}

func (vu *cardUsecases) AddCard(ctx context.Context, userID string, card entities.Card) error {
	return vu.cardRepo.Create(ctx, userID, card)
}

func (vu *cardUsecases) GetCard(ctx context.Context, userID string, meta string) (entities.Card, error) {
	return vu.cardRepo.GetOne(ctx, userID, meta)
}

func (vu *cardUsecases) DeleteCard(ctx context.Context, userID string, meta string) error {
	return vu.cardRepo.Delete(ctx, userID, meta)
}

func (vu *cardUsecases) ListCard(ctx context.Context, userID string) ([]entities.Card, error) {
	return vu.cardRepo.GetList(ctx, userID)
}

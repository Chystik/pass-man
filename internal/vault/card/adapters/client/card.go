package adapters

import (
	"context"
	"errors"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/vault/card/adapters/converter"
	"github.com/Chystik/pass-man/internal/vault/card/entities"
)

type CardAPIClient interface {
	AddCard(ctx context.Context, card entities.Card) error
	GetCard(ctx context.Context, meta string) (entities.Card, error)
	ListCard(ctx context.Context) ([]entities.Card, error)
	DeleteCard(ctx context.Context, meta string) error
}

type cardAPIClient struct {
	card pb.CardServiceClient
	CardAPIClient
}

func NewCardAPIClient(card pb.CardServiceClient) *cardAPIClient {
	return &cardAPIClient{
		card: card,
	}
}

func (ca *cardAPIClient) AddCard(ctx context.Context, card entities.Card) error {
	req := &pb.AddCardRequest{
		Card: converter.FromDomainCard(card),
	}

	res, err := ca.card.AddCard(ctx, req)
	if err != nil {
		return nil
	}

	if res.Error != nil {
		return errors.New(res.Error.String())
	}

	return nil
}

func (ca *cardAPIClient) GetCard(ctx context.Context, meta string) (entities.Card, error) {
	c := entities.Card{}

	req := &pb.GetCardRequest{
		Meta: meta,
	}

	res, err := ca.card.GetCard(ctx, req)
	if err != nil {
		return c, err
	}

	if res.Error != nil {
		return c, errors.New(res.Error.String())
	}

	return converter.ToDomainCard(res.Card), nil
}

func (ca *cardAPIClient) ListCard(ctx context.Context) ([]entities.Card, error) {
	c := []entities.Card{}
	req := &pb.ListCardRequest{}

	res, err := ca.card.ListCard(ctx, req)
	if err != nil {
		return c, err
	}

	if res.Error != nil {
		return c, errors.New(res.Error.String())
	}

	return converter.ToDomainCards(res.CardList), nil
}

func (ca *cardAPIClient) DeleteCard(ctx context.Context, meta string) error {
	req := &pb.DeleteCardRequest{
		Meta: meta,
	}

	res, err := ca.card.DeleteCard(ctx, req)
	if err != nil {
		return err
	}

	if res.Error != nil {
		return errors.New(res.Error.String())
	}

	return nil
}

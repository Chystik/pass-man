package adapters

import (
	"context"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/vault/adapters/converter"
	"github.com/Chystik/pass-man/internal/vault/usecases"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type cardHandlers struct {
	usecases usecases.CardUsecases
	pb.UnimplementedCardServiceServer
}

func NewCardHandlers(cu usecases.CardUsecases) *cardHandlers {
	return &cardHandlers{
		usecases: cu,
	}
}

func (ch *cardHandlers) AddCard(ctx context.Context, c *pb.AddCardRequest) (*pb.AddCardResponse, error) {
	var response pb.AddCardResponse

	userID, err := getUserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	card := converter.ToDomainCard(c.Card)

	err = ch.usecases.AddCard(ctx, userID, card)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "add card error: %s", err.Error())
	}

	return &response, nil
}

func (ch *cardHandlers) GetCard(ctx context.Context, c *pb.GetCardRequest) (*pb.GetCardResponse, error) {
	var response pb.GetCardResponse

	userID, err := getUserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	card, err := ch.usecases.GetCard(ctx, userID, c.Meta)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get card error: %s", err.Error())
	}

	response.Card = converter.FromDomainCard(card)

	return &response, err
}

func (ch *cardHandlers) ListCard(ctx context.Context, c *pb.ListCardRequest) (*pb.ListCardResponse, error) {
	var response pb.ListCardResponse

	userID, err := getUserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	cardList, err := ch.usecases.ListCard(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list card error: %s", err.Error())
	}

	response.CardList = converter.FromDomainCards(cardList)

	return &response, nil
}

func (ch *cardHandlers) DeleteCard(ctx context.Context, c *pb.DeleteCardRequest) (*pb.DeleteCardResponse, error) {
	var response pb.DeleteCardResponse

	userID, err := getUserFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	err = ch.usecases.DeleteCard(ctx, userID, c.Meta)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "delete card error: %s", err.Error())
	}

	return &response, nil
}

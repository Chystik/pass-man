package adapters

import (
	"context"
	"errors"
	"testing"

	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	mocks "github.com/Chystik/pass-man/internal/infrastructure/grpc/mock"
	"github.com/Chystik/pass-man/internal/vault/card/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddCard_ReturnsNoErr(t *testing.T) {
	t.Parallel()

	mk, cl := newCardAPIClientMock()
	ctx := context.Background()
	resp := &pb.AddCardResponse{}

	mk.card.EXPECT().AddCard(ctx, mock.Anything).Return(resp, nil)
	err := cl.AddCard(ctx, entities.Card{})
	assert.NoError(t, err)
}

func TestAddCard_WhenClientAddCardReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newCardAPIClientMock()
	ctx := context.Background()

	mk.card.EXPECT().AddCard(ctx, mock.Anything).Return(nil, errors.New("err"))
	err := cl.AddCard(ctx, entities.Card{})

	assert.Error(t, err)
}

func TestAddCard_WhenRespReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newCardAPIClientMock()
	ctx := context.Background()
	resp := &pb.AddCardResponse{
		Error: &pb.Error{},
	}

	mk.card.EXPECT().AddCard(ctx, mock.Anything).Return(resp, nil)
	err := cl.AddCard(ctx, entities.Card{})

	assert.Error(t, err)
}

func TestGetCard_ReturnsNoErr(t *testing.T) {
	t.Parallel()

	mk, cl := newCardAPIClientMock()
	ctx := context.Background()
	resp := &pb.GetCardResponse{
		Card: &pb.Card{},
	}

	mk.card.EXPECT().GetCard(ctx, mock.Anything).Return(resp, nil)
	_, err := cl.GetCard(ctx, "meta")

	assert.NoError(t, err)
}

func TestGetCard_WhenClientGetCardReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newCardAPIClientMock()
	ctx := context.Background()

	mk.card.EXPECT().GetCard(ctx, mock.Anything).Return(nil, errors.New("err"))
	_, err := cl.GetCard(ctx, "meta")

	assert.Error(t, err)
}

func TestGetCard_WhenRespReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newCardAPIClientMock()
	ctx := context.Background()
	resp := &pb.GetCardResponse{
		Error: &pb.Error{},
	}

	mk.card.EXPECT().GetCard(ctx, mock.Anything).Return(resp, nil)
	_, err := cl.GetCard(ctx, "meta")

	assert.Error(t, err)
}

func TestListCard_ReturnsNoErr(t *testing.T) {
	t.Parallel()

	mk, cl := newCardAPIClientMock()
	ctx := context.Background()
	resp := &pb.ListCardResponse{
		CardList: []*pb.Card{},
	}

	mk.card.EXPECT().ListCard(ctx, mock.Anything).Return(resp, nil)
	_, err := cl.ListCard(ctx)

	assert.NoError(t, err)
}

func TestListCard_WhenClientListCardReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newCardAPIClientMock()
	ctx := context.Background()

	mk.card.EXPECT().ListCard(ctx, mock.Anything).Return(nil, errors.New("err"))
	_, err := cl.ListCard(ctx)

	assert.Error(t, err)
}

func TestListCard_WhenRespReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newCardAPIClientMock()
	ctx := context.Background()
	resp := &pb.ListCardResponse{
		Error: &pb.Error{},
	}

	mk.card.EXPECT().ListCard(ctx, mock.Anything).Return(resp, nil)
	_, err := cl.ListCard(ctx)

	assert.Error(t, err)
}

func TestDeleteCard_ReturnsNoErr(t *testing.T) {
	t.Parallel()

	mk, cl := newCardAPIClientMock()
	ctx := context.Background()
	resp := &pb.DeleteCardResponse{}

	mk.card.EXPECT().DeleteCard(ctx, mock.Anything).Return(resp, nil)
	err := cl.DeleteCard(ctx, "meta")

	assert.NoError(t, err)
}

func TestDeleteCard_WhenClientDeleteCardReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newCardAPIClientMock()
	ctx := context.Background()

	mk.card.EXPECT().DeleteCard(ctx, mock.Anything).Return(nil, errors.New("err"))
	err := cl.DeleteCard(ctx, "meta")

	assert.Error(t, err)
}

func TestDeleteCard_WhenRespReturnsErr(t *testing.T) {
	t.Parallel()

	mk, cl := newCardAPIClientMock()
	ctx := context.Background()
	resp := &pb.DeleteCardResponse{
		Error: &pb.Error{},
	}

	mk.card.EXPECT().DeleteCard(ctx, mock.Anything).Return(resp, nil)
	err := cl.DeleteCard(ctx, "meta")

	assert.Error(t, err)
}

type cardAPIClientMock struct {
	card *mocks.CardServiceClient
}

func newCardAPIClientMock() (*cardAPIClientMock, *cardAPIClient) {
	m := &cardAPIClientMock{
		card: &mocks.CardServiceClient{},
	}
	return m, NewCardAPIClient(m.card)
}

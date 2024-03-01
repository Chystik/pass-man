package converter

import (
	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/vault/card/entities"
)

func ToDomainCard(c *pb.Card) entities.Card {
	res := entities.Card{
		Meta:   c.Meta,
		Number: c.Number,
		Holder: c.Holder,
		CVV:    int(c.Cvv),
	}

	res.ValidThru.Parse(c.ValidThru)

	return res
}

func ToDomainCards(c []*pb.Card) []entities.Card {
	res := make([]entities.Card, len(c))

	for i := range c {
		res[i] = ToDomainCard(c[i])
	}

	return res
}

func FromDomainCard(c entities.Card) *pb.Card {
	return &pb.Card{
		Meta:      c.Meta,
		Number:    c.Number,
		ValidThru: c.ValidThru.String(),
		Holder:    c.Holder,
		Cvv:       uint32(c.CVV),
	}
}

func FromDomainCards(c []entities.Card) []*pb.Card {
	res := make([]*pb.Card, len(c))

	for i := range c {
		res[i] = FromDomainCard(c[i])
	}

	return res
}

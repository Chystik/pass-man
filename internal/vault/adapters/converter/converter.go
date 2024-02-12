package converter

import (
	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/vault/entities"
)

func ToDomainPassword(p *pb.Password) entities.Password {
	return entities.Password{
		Meta:     p.Meta,
		Username: p.Username,
		Password: p.Password,
	}
}

func ToDomainPasswords(p []*pb.Password) []entities.Password {
	res := make([]entities.Password, len(p))

	for i := range p {
		res[i] = ToDomainPassword(p[i])
	}

	return res
}

func FromDomainPassword(p entities.Password) *pb.Password {
	return &pb.Password{
		Meta:     p.Meta,
		Username: p.Username,
		Password: p.Password,
	}
}

func FromDomainPasswords(p []entities.Password) []*pb.Password {
	res := make([]*pb.Password, len(p))

	for i := range p {
		res[i] = FromDomainPassword(p[i])
	}

	return res
}

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

func ToDomainFile(f *pb.File) *entities.File {
	return &entities.File{
		ID:   f.Id,
		Meta: f.Meta,
		Name: f.FullName,
	}
}

func ToDomainFiles(f []*pb.File) []*entities.File {
	res := make([]*entities.File, len(f))

	for i := range f {
		res[i] = ToDomainFile(f[i])
	}

	return res
}

func FromDomainFile(f *entities.File) *pb.File {
	return &pb.File{
		Id:       f.ID,
		Meta:     f.Meta,
		FullName: f.Name,
	}
}

func FromDomainFiles(f []*entities.File) []*pb.File {
	res := make([]*pb.File, len(f))

	for i := range f {
		res[i] = FromDomainFile(f[i])
	}

	return res
}

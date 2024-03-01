package converter

import (
	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/vault/password/entities"
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

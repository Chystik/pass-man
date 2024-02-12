package converter

import (
	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/vault/note/entities"
)

func ToDomainNote(n *pb.Note) entities.Note {
	return entities.Note{
		Meta: n.Meta,
		Note: n.Note,
	}
}

func ToDomainNotes(n []*pb.Note) []entities.Note {
	res := make([]entities.Note, len(n))

	for i := range n {
		res[i] = ToDomainNote(n[i])
	}

	return res
}

func FromDomainNote(n entities.Note) *pb.Note {
	return &pb.Note{
		Meta: n.Meta,
		Note: n.Note,
	}
}

func FromDomainNotes(n []entities.Note) []*pb.Note {
	res := make([]*pb.Note, len(n))

	for i := range n {
		res[i] = FromDomainNote(n[i])
	}

	return res
}

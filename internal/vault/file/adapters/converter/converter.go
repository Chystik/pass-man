package converter

import (
	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/vault/file/entities"
)

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

package usecases

import (
	"context"

	"github.com/Chystik/pass-man/internal/vault/file/entities"
)

type FileUsecases interface {
	Upload(ctx context.Context, userID string, file *entities.File) (int, error)
	Download(ctx context.Context, userID string, file *entities.File) (int, error)
	Delete(ctx context.Context, userID string, file *entities.File) error
	ListFiles(ctx context.Context, userID string) ([]*entities.File, error)
}

type fileUsecases struct {
	fileRepo FileRepository
}

func NewFileUsecases(f FileRepository) *fileUsecases {
	return &fileUsecases{
		fileRepo: f,
	}
}

func (fu *fileUsecases) Upload(ctx context.Context, userID string, file *entities.File) (int, error) {
	return fu.fileRepo.Create(ctx, userID, file)
}

func (fu *fileUsecases) Download(ctx context.Context, userID string, file *entities.File) (int, error) {
	return fu.fileRepo.GetOne(ctx, userID, file)
}

func (fu *fileUsecases) Delete(ctx context.Context, userID string, file *entities.File) error {
	return fu.fileRepo.Delete(ctx, userID, file)
}

func (fu *fileUsecases) ListFiles(ctx context.Context, userID string) ([]*entities.File, error) {
	return fu.fileRepo.GetList(ctx, userID)
}

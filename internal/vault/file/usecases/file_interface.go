package usecases

import (
	"context"

	"github.com/Chystik/pass-man/internal/vault/file/entities"
)

type FileRepository interface {
	Create(ctx context.Context, userID string, file *entities.File) (int, error)
	GetOne(ctx context.Context, userID string, file *entities.File) (int, error)
	Delete(ctx context.Context, userID string, file *entities.File) error
	GetList(ctx context.Context, userID string) ([]*entities.File, error)
}

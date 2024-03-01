package usecases

import (
	"context"

	"github.com/Chystik/pass-man/internal/vault/note/entities"
)

type NoteRepository interface {
	Create(ctx context.Context, userID string, note entities.Note) error
	GetOne(ctx context.Context, userID string, meta string) (entities.Note, error)
	Delete(ctx context.Context, userID string, meta string) error
	GetList(ctx context.Context, userID string) ([]entities.Note, error)
}

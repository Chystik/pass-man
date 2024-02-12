package usecases

import (
	"context"

	"github.com/Chystik/pass-man/internal/vault/note/entities"
)

type NoteUsecases interface {
	AddNote(ctx context.Context, userID string, note entities.Note) error
	GetNote(ctx context.Context, userID string, meta string) (entities.Note, error)
	DeleteNote(ctx context.Context, userID string, meta string) error
	ListNote(ctx context.Context, userID string) ([]entities.Note, error)
}

type noteUsecases struct {
	noteRepo NoteRepository
}

func NewNoteRepository(n NoteRepository) *noteUsecases {
	return &noteUsecases{
		noteRepo: n,
	}
}

func (nr *noteUsecases) AddNote(ctx context.Context, userID string, note entities.Note) error {
	return nr.noteRepo.Create(ctx, userID, note)
}

func (nr *noteUsecases) GetNote(ctx context.Context, userID string, meta string) (entities.Note, error) {
	return nr.noteRepo.GetOne(ctx, userID, meta)
}

func (nr *noteUsecases) DeleteNote(ctx context.Context, userID string, meta string) error {
	return nr.noteRepo.Delete(ctx, userID, meta)
}

func (nr *noteUsecases) ListNote(ctx context.Context, userID string) ([]entities.Note, error) {
	return nr.noteRepo.GetList(ctx, userID)
}

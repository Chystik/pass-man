package usecases

import (
	"context"

	"github.com/Chystik/pass-man/internal/vault/entities"
)

type PasswordAPIClient interface {
	AddPassword(ctx context.Context, password entities.Password) error
	GetPassword(ctx context.Context, meta string) (entities.Password, error)
	ListPassword(ctx context.Context) ([]entities.Password, error)
	DeletePassword(ctx context.Context, meta string) error
}

type CardAPIClient interface {
	AddCard(ctx context.Context, card entities.Card) error
	GetCard(ctx context.Context, meta string) (entities.Card, error)
	ListCard(ctx context.Context) ([]entities.Card, error)
	DeleteCard(ctx context.Context, meta string) error
}

type NoteAPIClient interface {
	AddNote(ctx context.Context, password entities.Password) error
	GetNote(ctx context.Context) (entities.Password, error)
	ListNote(ctx context.Context) ([]entities.Password, error)
}

type FileAPIClient interface {
	UploadFile(ctx context.Context, file *entities.File, filePath string) error
	DownloadFile(ctx context.Context, file *entities.File, filePath string) error
	ListFiles(ctx context.Context) ([]*entities.File, error)
	DeleteFile(ctx context.Context, file *entities.File) error
}

type VaultAPICliet interface {
	PasswordAPIClient
	CardAPIClient
	FileAPIClient
}

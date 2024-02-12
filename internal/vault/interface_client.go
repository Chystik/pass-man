package vault

import (
	"context"

	card "github.com/Chystik/pass-man/internal/vault/card/entities"
	file "github.com/Chystik/pass-man/internal/vault/file/entities"
	note "github.com/Chystik/pass-man/internal/vault/note/entities"
	pass "github.com/Chystik/pass-man/internal/vault/password/entities"
)

type PasswordAPIClient interface {
	AddPassword(ctx context.Context, password pass.Password) error
	GetPassword(ctx context.Context, meta string) (pass.Password, error)
	ListPassword(ctx context.Context) ([]pass.Password, error)
	DeletePassword(ctx context.Context, meta string) error
}

type CardAPIClient interface {
	AddCard(ctx context.Context, card card.Card) error
	GetCard(ctx context.Context, meta string) (card.Card, error)
	ListCard(ctx context.Context) ([]card.Card, error)
	DeleteCard(ctx context.Context, meta string) error
}

type NoteAPIClient interface {
	AddNote(ctx context.Context, note note.Note) error
	GetNote(ctx context.Context, meta string) (note.Note, error)
	ListNote(ctx context.Context) ([]note.Note, error)
	DeleteNote(ctx context.Context, meta string) error
}

type FileAPIClient interface {
	UploadFile(ctx context.Context, file *file.File, filePath string) error
	DownloadFile(ctx context.Context, file *file.File, filePath string) error
	ListFiles(ctx context.Context) ([]*file.File, error)
	DeleteFile(ctx context.Context, file *file.File) error
}

type VaultAPICliet interface {
	PasswordAPIClient
	CardAPIClient
	FileAPIClient
	NoteAPIClient
}

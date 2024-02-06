package usecases

import (
	"context"
	"io"

	"github.com/Chystik/pass-man/internal/vault/entities"
)

type PasswordRepository interface {
	AddPassword(ctx context.Context, password entities.Password) error
	GetPassword(ctx context.Context) (entities.Password, error)
	DeletePassword(ctx context.Context, password entities.Password) error
	ListPassword(ctx context.Context) ([]entities.Password, error)
}

type CardRepository interface {
	AddCard(ctx context.Context, card entities.Card) error
	GetCard(ctx context.Context) (entities.Card, error)
	DeleteCard(ctx context.Context, car entities.Password) error
	ListCard(ctx context.Context) ([]entities.Card, error)
}

type VaultRepository interface {
	PasswordRepository
	CardRepository
}

type VaultCryptor interface {
	Encrypt(dst io.Writer) (int, error)
	Decrypt(src io.Reader) (int, error)
}

type VaultKeyStore interface {
	Lock(login string) error
	Unlock(login string, key []byte) error
	GetKey(login string) ([]byte, error)
}

type PasswordUsecase interface {
	AddPassword(ctx context.Context, password entities.Password) error
	GetPassword(ctx context.Context) (entities.Password, error)
	DeletePassword(ctx context.Context, password entities.Password) error
	ListPassword(ctx context.Context) ([]entities.Password, error)
}

type VaultUsecase interface {
	PasswordUsecase
}

type PasswordAPIClient interface {
	AddPassword(ctx context.Context, password entities.Password) error
	GetPassword(ctx context.Context) (entities.Password, error)
	DeletePassword(ctx context.Context, password entities.Password) error
	ListPassword(ctx context.Context) ([]entities.Password, error)
}

type CardAPIClient interface {
	AddCard(ctx context.Context, password entities.Password) error
	GetCard(ctx context.Context) (entities.Password, error)
	ListCard(ctx context.Context) ([]entities.Password, error)
}

type NoteAPIClient interface {
	AddNote(ctx context.Context, password entities.Password) error
	GetNote(ctx context.Context) (entities.Password, error)
	ListNote(ctx context.Context) ([]entities.Password, error)
}

type FilePIClient interface {
	AddFile(ctx context.Context, password entities.Password) error
	GetFile(ctx context.Context) (entities.Password, error)
	ListFile(ctx context.Context) ([]entities.Password, error)
}

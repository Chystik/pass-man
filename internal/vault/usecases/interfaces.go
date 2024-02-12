package usecases

import (
	"context"
	"io"

	"github.com/Chystik/pass-man/internal/vault/entities"
)

type VaultCryptor interface {
	Encrypt(in io.Reader, out io.Writer, userID string) (int, error)
	Decrypt(in io.Reader, out io.Writer, userID string) (int, error)
}

type VaultKeyStore interface {
	Lock(login string) error
	Unlock(login string, key []byte) error
	GetKey(login string) ([]byte, error)
}

type PasswordRepository interface {
	Create(ctx context.Context, userID string, password entities.Password) error
	GetOne(ctx context.Context, userID string, meta string) (entities.Password, error)
	Delete(ctx context.Context, userID string, meta string) error
	GetList(ctx context.Context, userID string) ([]entities.Password, error)
}

type CardRepository interface {
	Create(ctx context.Context, userID string, card entities.Card) error
	GetOne(ctx context.Context, userID string, meta string) (entities.Card, error)
	Delete(ctx context.Context, userID string, meta string) error
	GetList(ctx context.Context, userID string) ([]entities.Card, error)
}

type FileRepository interface {
	Create(ctx context.Context, userID string, file *entities.File) (int, error)
	GetOne(ctx context.Context, userID string, file *entities.File) (int, error)
	Delete(ctx context.Context, userID string, file *entities.File) error
	GetList(ctx context.Context, userID string) ([]*entities.File, error)
}

type PasswordUsecases interface {
	AddPassword(ctx context.Context, userID string, password entities.Password) error
	GetPassword(ctx context.Context, userID string, meta string) (entities.Password, error)
	DeletePassword(ctx context.Context, userID string, meta string) error
	ListPassword(ctx context.Context, userID string) ([]entities.Password, error)
}

type CardUsecases interface {
	AddCard(ctx context.Context, userID string, card entities.Card) error
	GetCard(ctx context.Context, userID string, meta string) (entities.Card, error)
	DeleteCard(ctx context.Context, userID string, meta string) error
	ListCard(ctx context.Context, userID string) ([]entities.Card, error)
}

type FileUsecases interface {
	Upload(ctx context.Context, userID string, file *entities.File) (int, error)
	Download(ctx context.Context, userID string, file *entities.File) (int, error)
	Delete(ctx context.Context, userID string, file *entities.File) error
	ListFiles(ctx context.Context, userID string) ([]*entities.File, error)
}

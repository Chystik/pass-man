package repository

import (
	"context"

	"github.com/Chystik/pass-man/internal/vault/entities"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type passwordRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewPasswordRepository(db *sqlx.DB, log *zap.Logger) *passwordRepository {
	return &passwordRepository{
		db:  db,
		log: log,
	}
}

func (pr *passwordRepository) AddPassword(ctx context.Context, password entities.Password) error {
	panic("IMPLEMENT ME")
}

func (pr *passwordRepository) GetPassword(ctx context.Context) (entities.Password, error) {
	panic("IMPLEMENT ME")
}

func (pr *passwordRepository) DeletePassword(ctx context.Context, password entities.Password) error {
	panic("IMPLEMENT ME")
}

func (pr *passwordRepository) ListPassword(ctx context.Context) ([]entities.Password, error) {
	panic("IMPLEMENT ME")
}

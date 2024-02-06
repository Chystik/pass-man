package repository

import (
	"context"
	"fmt"

	appErr "github.com/Chystik/pass-man/internal/error/entities"
	"github.com/Chystik/pass-man/internal/user/entities"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type userRepository struct {
	db  *sqlx.DB
	log *zap.Logger
}

func NewUserRepository(db *sqlx.DB, log *zap.Logger) *userRepository {
	return &userRepository{
		db:  db,
		log: log,
	}
}

func (ur *userRepository) Create(ctx context.Context, user entities.User) error {
	query := `
			INSERT INTO	passman.user (login, password, vault_key)
			VALUES ($1, $2, $3)`

	ur.log.Debug("UserRepository.Create", zap.String("query", query))

	_, err := ur.db.ExecContext(ctx, query, user.Login, user.HashedPassword, user.EncryptedVaultKey)
	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if !ok {
			return err
		} else if pgErr.Code == "23505" { // login exists: duplicate key value violates unique constraint
			fmt.Println("true1")
			return &appErr.AppError{Op: "userRepository.Create", Code: appErr.ErrExists, Message: fmt.Sprintf("user %s already exists", user.Login)}
		}
		ur.log.Error(err.Error())
		return err
	}

	return nil
}

func (ur *userRepository) Get(ctx context.Context, login string) (entities.User, error) {
	var u entities.User

	query := `
			SELECT login, password, vault_key
			FROM passman.user
			WHERE login = $1`

	err := ur.db.GetContext(ctx, &u, query, login)
	if err != nil {
		return u, err
	}

	return u, nil
}

func (ur *userRepository) Update(ctx context.Context, user entities.User) error {
	return nil
}

func (ur *userRepository) Delete(ctx context.Context, user entities.User) error {
	return nil
}

package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	appErr "github.com/Chystik/pass-man/internal/error/entities"
	"github.com/Chystik/pass-man/internal/vault"
	"github.com/Chystik/pass-man/internal/vault/password/entities"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type passwordRepository struct {
	db      *sqlx.DB
	log     *zap.Logger
	cryptor vault.VaultCryptor
}

func NewPasswordRepository(db *sqlx.DB, log *zap.Logger, cryptor vault.VaultCryptor) *passwordRepository {
	return &passwordRepository{
		db:      db,
		log:     log,
		cryptor: cryptor,
	}
}

func (pr *passwordRepository) Create(ctx context.Context, userID string, password entities.Password) error {
	p, err := pr.fromDomainPassword(password, userID)
	if err != nil {
		return err
	}

	query := `
			INSERT INTO	passman.password (user_id, meta, username, password)
			VALUES ($1, $2, $3, $4)`

	_, err = pr.db.ExecContext(ctx, query, p.UserID, p.Meta, p.Username, p.Password)
	if err != nil {
		pr.log.Error(err.Error())
		return &appErr.AppError{Op: "PasswordRepository.Create", Message: err.Error()}
	}

	return nil
}

func (pr *passwordRepository) GetOne(ctx context.Context, userID string, meta string) (entities.Password, error) {
	p := dsPassword{}
	res := entities.Password{}

	query := `
			SELECT meta, username, password
			FROM passman.password
			WHERE user_id = $1 AND meta = $2`

	err := pr.db.GetContext(ctx, &p, query, userID, meta)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, &appErr.AppError{Op: "PasswordRepository.GetOne", Code: appErr.ErrNotFound, Message: fmt.Sprintf("password with descriprion %s not found", meta)}
		}
		pr.log.Error(err.Error())
		return res, err
	}

	res, err = pr.toDomainPassword(p, userID)
	if err != nil {
		pr.log.Error(err.Error())
		return res, err
	}

	return res, nil
}

func (pr *passwordRepository) GetList(ctx context.Context, userID string) ([]entities.Password, error) {
	p := []dsPassword{}
	res := []entities.Password{}

	query := `
			SELECT meta, username, password
			FROM passman.password
			WHERE user_id = $1`

	err := pr.db.SelectContext(ctx, &p, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, &appErr.AppError{Op: "PasswordRepository.GetList", Code: appErr.ErrEmptyList, Message: "no passwords in vault"}
		}
		pr.log.Error(err.Error())
		return res, err
	}

	if len(p) == 0 {
		return res, &appErr.AppError{Code: appErr.ErrEmptyList}
	}

	res, err = pr.toDomainPasswords(p, userID)
	if err != nil {
		pr.log.Error(err.Error())
		return res, err
	}

	return res, nil
}

func (pr *passwordRepository) Delete(ctx context.Context, userID string, meta string) error {
	query := `
			DELETE FROM passman.password
			WHERE user_id = $1 AND meta = $2`

	res, err := pr.db.ExecContext(ctx, query, userID, meta)
	if err != nil {
		pr.log.Error(err.Error())
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		pr.log.Error(err.Error())
		return err
	}

	if n == 0 {
		return &appErr.AppError{Code: appErr.ErrNotFound}
	}

	return nil
}

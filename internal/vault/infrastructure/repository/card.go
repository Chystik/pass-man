package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	appErr "github.com/Chystik/pass-man/internal/error/entities"
	"github.com/Chystik/pass-man/internal/vault/entities"
	"github.com/Chystik/pass-man/internal/vault/usecases"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type cardRepository struct {
	db      *sqlx.DB
	log     *zap.Logger
	cryptor usecases.VaultCryptor
}

func NewCardRepository(db *sqlx.DB, log *zap.Logger, cryptor usecases.VaultCryptor) *cardRepository {
	return &cardRepository{
		db:      db,
		log:     log,
		cryptor: cryptor,
	}
}

func (cr *cardRepository) Create(ctx context.Context, userID string, card entities.Card) error {
	c, err := cr.fromDomainCard(card, userID)
	if err != nil {
		return err
	}

	query := `
			INSERT INTO	passman.card (user_id, meta, number, valid_thru, holder, cvv)
			VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = cr.db.ExecContext(ctx, query, c.UserID, c.Meta, c.Number, c.ValidThru, c.Holder, c.CVV)
	if err != nil {
		cr.log.Error(err.Error())
		return &appErr.AppError{Op: "CardRepository.Create", Message: err.Error()}
	}

	return nil
}

func (cr *cardRepository) GetOne(ctx context.Context, userID string, meta string) (entities.Card, error) {
	c := dsCard{}
	res := entities.Card{}

	query := `
			SELECT meta, number, valid_thru, holder, cvv
			FROM passman.card
			WHERE user_id = $1 AND meta = $2`

	err := cr.db.GetContext(ctx, &c, query, userID, meta)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, &appErr.AppError{Op: "CardRepository.GetOne", Code: appErr.ErrNotFound, Message: fmt.Sprintf("password with descriprion %s not found", meta)}
		}
		cr.log.Error(err.Error())
		return res, err
	}

	res, err = cr.toDomainCard(c, userID)
	if err != nil {
		cr.log.Error(err.Error())
		return res, err
	}

	return res, nil
}

func (cr *cardRepository) GetList(ctx context.Context, userID string) ([]entities.Card, error) {
	c := []dsCard{}
	res := []entities.Card{}

	query := `
			SELECT meta, number, valid_thru, holder, cvv
			FROM passman.card
			WHERE user_id = $1`

	err := cr.db.SelectContext(ctx, &c, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, &appErr.AppError{Op: "CardRepository.GetList", Code: appErr.ErrEmptyList, Message: "no passwords in vault"}
		}
		cr.log.Error(err.Error())
		return res, err
	}

	if len(c) == 0 {
		return res, &appErr.AppError{Code: appErr.ErrEmptyList}
	}

	res, err = cr.toDomainCards(c, userID)
	if err != nil {
		cr.log.Error(err.Error())
		return res, err
	}

	return res, nil
}

func (cr *cardRepository) Delete(ctx context.Context, userID string, meta string) error {
	query := `
			DELETE FROM passman.card
			WHERE user_id = $1 AND meta = $2`

	res, err := cr.db.ExecContext(ctx, query, userID, meta)
	if err != nil {
		cr.log.Error(err.Error())
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		cr.log.Error(err.Error())
		return err
	}

	if n == 0 {
		return &appErr.AppError{Code: appErr.ErrNotFound}
	}

	return nil
}

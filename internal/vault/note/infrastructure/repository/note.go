package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	appErr "github.com/Chystik/pass-man/internal/error/entities"
	"github.com/Chystik/pass-man/internal/vault"
	"github.com/Chystik/pass-man/internal/vault/note/entities"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type noteRepository struct {
	db      *sqlx.DB
	log     *zap.Logger
	cryptor vault.VaultCryptor
}

func NewNoteRepository(db *sqlx.DB, log *zap.Logger, cryptor vault.VaultCryptor) *noteRepository {
	return &noteRepository{
		db:      db,
		log:     log,
		cryptor: cryptor,
	}
}

func (nr *noteRepository) Create(ctx context.Context, userID string, note entities.Note) error {
	n, err := nr.fromDomainNote(note, userID)
	if err != nil {
		return err
	}

	query := `
			INSERT INTO	passman.note (user_id, meta, note)
			VALUES ($1, $2, $3)`

	_, err = nr.db.ExecContext(ctx, query, n.UserID, n.Meta, n.Note)
	if err != nil {
		nr.log.Error(err.Error())
		return &appErr.AppError{Op: "NoteRepository.Create", Message: err.Error()}
	}

	return nil
}

func (nr *noteRepository) GetOne(ctx context.Context, userID string, meta string) (entities.Note, error) {
	n := dsNote{}
	res := entities.Note{}

	query := `
			SELECT meta, note
			FROM passman.note
			WHERE user_id = $1 AND meta = $2`

	err := nr.db.GetContext(ctx, &n, query, userID, meta)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, &appErr.AppError{Op: "PasswordRepository.GetOne", Code: appErr.ErrNotFound, Message: fmt.Sprintf("password with descriprion %s not found", meta)}
		}
		nr.log.Error(err.Error())
		return res, err
	}

	res, err = nr.toDomainNote(n, userID)
	if err != nil {
		nr.log.Error(err.Error())
		return res, err
	}

	return res, nil
}

func (nr *noteRepository) Delete(ctx context.Context, userID string, meta string) error {
	query := `
		DELETE FROM passman.note
		WHERE user_id = $1 AND meta = $2`

	res, err := nr.db.ExecContext(ctx, query, userID, meta)
	if err != nil {
		nr.log.Error(err.Error())
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		nr.log.Error(err.Error())
		return err
	}

	if n == 0 {
		return &appErr.AppError{Code: appErr.ErrNotFound}
	}

	return nil
}

func (nr *noteRepository) GetList(ctx context.Context, userID string) ([]entities.Note, error) {
	n := []dsNote{}
	res := []entities.Note{}

	query := `
			SELECT meta, note
			FROM passman.note
			WHERE user_id = $1`

	err := nr.db.SelectContext(ctx, &n, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, &appErr.AppError{Op: "NoteRepository.GetList", Code: appErr.ErrEmptyList, Message: "no passwords in vault"}
		}
		nr.log.Error(err.Error())
		return res, err
	}

	if len(n) == 0 {
		return res, &appErr.AppError{Code: appErr.ErrEmptyList}
	}

	res, err = nr.toDomainNotes(n, userID)
	if err != nil {
		nr.log.Error(err.Error())
		return res, err
	}

	return res, nil
}

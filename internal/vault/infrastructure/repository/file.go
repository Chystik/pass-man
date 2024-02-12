package repository

import (
	"context"
	"database/sql"
	"errors"

	appErr "github.com/Chystik/pass-man/internal/error/entities"
	"github.com/Chystik/pass-man/internal/vault/entities"
	"github.com/Chystik/pass-man/internal/vault/usecases"

	"github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type fileRepository struct {
	db      *sqlx.DB
	conn    *pgx.Conn
	log     *zap.Logger
	cryptor usecases.VaultCryptor
}

func NewFileRepository(db *sqlx.DB, c *pgx.Conn, log *zap.Logger, cryptor usecases.VaultCryptor) *fileRepository {
	return &fileRepository{
		db:      db,
		conn:    c,
		log:     log,
		cryptor: cryptor,
	}
}

func (fr *fileRepository) Create(ctx context.Context, userID string, file *entities.File) (int, error) {
	f, err := fr.fromDomainFile(file, userID)
	if err != nil {
		fr.log.Error(err.Error())
		return 0, err
	}

	tx, err := fr.conn.Begin(ctx)
	if err != nil {
		fr.log.Error(err.Error())
		return 0, err
	}
	defer tx.Rollback(ctx)

	lobs := tx.LargeObjects()

	// Create a new Large Object.
	// We pass 0, so the DB can pick an available oid for us.
	oid, err := lobs.Create(ctx, 0)
	if err != nil {
		fr.log.Error(err.Error())
		return 0, err
	}

	query := `
			INSERT INTO passman.file (id, user_id, meta, full_name) 
			VALUES ($1, $2, $3, $4)`

	// Record the oid and filename in the files table
	_, err = tx.Exec(ctx, query, oid, userID, f.Meta, f.Name)
	if err != nil {
		fr.log.Error(err.Error())
		return 0, err
	}

	// Open the new Object for writing.
	obj, err := lobs.Open(ctx, oid, pgx.LargeObjectModeWrite)
	if err != nil {
		fr.log.Error(err.Error())
		return 0, err
	}

	// Copy the file stream to the Large Object stream using cryptor
	written, err := fr.cryptor.Encrypt(file.Data, obj, userID)
	if err != nil {
		fr.log.Error(err.Error())
		return written, err
	}

	err = tx.Commit(ctx)

	return written, err
}

func (fr *fileRepository) GetOne(ctx context.Context, userID string, file *entities.File) (int, error) {
	f := dsFile{}

	tx, err := fr.conn.Begin(ctx)
	if err != nil {
		fr.log.Error(err.Error())
		return 0, err
	}
	defer tx.Rollback(ctx)

	query := `
			SELECT id, user_id, meta, full_name
			FROM passman.file 
			WHERE user_id = $1 AND id = $2`

	err = fr.conn.QueryRow(ctx, query, userID, file.ID).Scan(&f.ID, &f.UserID, &f.Meta, &f.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, &appErr.AppError{Code: appErr.ErrNotFound}
		}
		fr.log.Error(err.Error())
		return 0, err
	}

	lobs := tx.LargeObjects()
	obj, err := lobs.Open(ctx, f.ID, pgx.LargeObjectModeRead)
	if err != nil {
		fr.log.Error(err.Error())
		return 0, err
	}

	read, err := fr.cryptor.Decrypt(obj, file.Data, userID)
	if err != nil {
		fr.log.Error(err.Error())
		return 0, err
	}

	return read, nil
}

func (fr *fileRepository) Delete(ctx context.Context, userID string, file *entities.File) error {
	query := `
			DELETE FROM passman.file
			WHERE user_id = $1 AND id =  $2`

	res, err := fr.db.ExecContext(ctx, query, userID, file.ID)
	if err != nil {
		fr.log.Error(err.Error())
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		fr.log.Error(err.Error())
		return err
	}

	if n == 0 {
		return &appErr.AppError{Code: appErr.ErrNotFound}
	}

	return nil
}

func (fr *fileRepository) GetList(ctx context.Context, userID string) ([]*entities.File, error) {
	f := []*dsFile{}
	res := []*entities.File{}

	query := `
			SELECT id, meta, full_name
			FROM passman.file
			WHERE user_id = $1`

	err := fr.db.SelectContext(ctx, &f, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, &appErr.AppError{Op: "CardRepository.GetList", Code: appErr.ErrEmptyList, Message: "no passwords in vault"}
		}
		fr.log.Error(err.Error())
		return res, err
	}

	if len(f) == 0 {
		return res, &appErr.AppError{Code: appErr.ErrEmptyList}
	}

	res, err = fr.toDomainFiles(f, userID)
	if err != nil {
		fr.log.Error(err.Error())
		return res, err
	}

	return res, nil
}

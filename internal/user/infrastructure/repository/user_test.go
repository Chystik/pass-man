package repository

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/Chystik/pass-man/internal/user/entities"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestCreate_WhenDBReturnsErr(t *testing.T) {
	t.Parallel()

	mockDB, mockSQL, mockLog := getUserRepoMocks(t)

	u := getTestUser()

	query := regexp.QuoteMeta(`
			INSERT INTO	passman.user (login, password, vault_key)
			VALUES ($1, $2, $3)`)
	repo := NewUserRepository(mockDB, mockLog)
	ctx := context.Background()

	mockSQL.ExpectExec(query).WithArgs(u.Login, u.HashedPassword, u.EncryptedVaultKey).WillReturnError(errors.New("err"))
	err := repo.Create(ctx, u)

	assert.Error(t, err)
}

func TestCreate_WhenDBReturnsPgconnErr(t *testing.T) {
	t.Parallel()

	mockDB, mockSQL, mockLog := getUserRepoMocks(t)

	u := getTestUser()

	query := regexp.QuoteMeta(`
			INSERT INTO	passman.user (login, password, vault_key)
			VALUES ($1, $2, $3)`)
	repo := NewUserRepository(mockDB, mockLog)
	ctx := context.Background()
	errPg := pgconn.PgError{}

	mockSQL.ExpectExec(query).WithArgs(u.Login, u.HashedPassword, u.EncryptedVaultKey).WillReturnError(&errPg)
	err := repo.Create(ctx, u)

	assert.Error(t, err)
}

func TestCreate_WhenDBReturnsPgconnDuplicadeErr(t *testing.T) {
	t.Parallel()

	mockDB, mockSQL, mockLog := getUserRepoMocks(t)

	u := getTestUser()

	query := regexp.QuoteMeta(`
			INSERT INTO	passman.user (login, password, vault_key)
			VALUES ($1, $2, $3)`)
	repo := NewUserRepository(mockDB, mockLog)
	ctx := context.Background()
	errPg := pgconn.PgError{Code: "23505"}

	mockSQL.ExpectExec(query).WithArgs(u.Login, u.HashedPassword, u.EncryptedVaultKey).WillReturnError(&errPg)
	err := repo.Create(ctx, u)

	assert.Error(t, err)
}

func TestCreate_ReturnsNoErr(t *testing.T) {
	t.Parallel()

	mockDB, mockSQL, mockLog := getUserRepoMocks(t)

	u := getTestUser()

	query := regexp.QuoteMeta(`
			INSERT INTO	passman.user (login, password, vault_key)
			VALUES ($1, $2, $3)`)
	repo := NewUserRepository(mockDB, mockLog)
	ctx := context.Background()

	mockSQL.ExpectExec(query).WithArgs(u.Login, u.HashedPassword, u.EncryptedVaultKey).WillReturnResult(sqlmock.NewResult(0, 0))
	err := repo.Create(ctx, u)

	assert.NoError(t, err)
}

func TestGet_ReturnsNoErr(t *testing.T) {
	t.Parallel()

	mockDB, mockSQL, mockLog := getUserRepoMocks(t)

	expected := getTestUser()
	query := regexp.QuoteMeta(`
			SELECT login, password, vault_key
			FROM passman.user
			WHERE login = $1`)
	repo := NewUserRepository(mockDB, mockLog)
	ctx := context.Background()
	u := getTestUser()

	mockSQL.ExpectQuery(query).WithArgs(expected.Login).WillReturnRows(sqlmock.NewRows([]string{
		"login",
		"password",
		"vault_key",
	}).AddRow(
		u.Login,
		u.HashedPassword,
		u.EncryptedVaultKey,
	))
	actual, err := repo.Get(ctx, u.Login)

	assert.Equal(t, expected, actual)
	assert.NoError(t, err)
}

func TestGet_WhenDBReturnsErr(t *testing.T) {
	t.Parallel()

	mockDB, mockSQL, mockLog := getUserRepoMocks(t)

	expected := entities.User{}
	query := regexp.QuoteMeta(`
			SELECT login, password, vault_key
			FROM passman.user
			WHERE login = $1`)
	repo := NewUserRepository(mockDB, mockLog)
	ctx := context.Background()
	u := getTestUser()

	mockSQL.ExpectQuery(query).WithArgs(expected.Login).WillReturnError(errors.New("err"))
	actual, err := repo.Get(ctx, u.Login)

	assert.Equal(t, expected, actual)
	assert.Error(t, err)
}

func TestGet_WhenDBReturnsNoRowsErr(t *testing.T) {
	t.Parallel()

	mockDB, mockSQL, mockLog := getUserRepoMocks(t)

	expected := entities.User{}
	query := regexp.QuoteMeta(`
			SELECT login, password, vault_key
			FROM passman.user
			WHERE login = $1`)
	repo := NewUserRepository(mockDB, mockLog)
	ctx := context.Background()
	u := getTestUser()
	errNoRows := sql.ErrNoRows

	mockSQL.ExpectQuery(query).WithArgs().WillReturnError(errNoRows)
	actual, err := repo.Get(ctx, u.Login)

	assert.Equal(t, expected, actual)
	assert.Error(t, err)
}

func getUserRepoMocks(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock, *zap.Logger) {
	mockDB, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		return nil, nil, nil
	}
	l := zaptest.NewLogger(t)
	return sqlx.NewDb(mockDB, "sqlmock"), sqlMock, l
}

func getTestUser() entities.User {
	return entities.User{
		Login: "test",
	}
}

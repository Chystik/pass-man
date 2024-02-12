package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const defaultPgPort uint16 = 5432

var (
	connStr            = "host=%s port=%d user=%s password=%s sslmode=%s"
	connStrDB          = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	logDatabaseCreated = "database %s created"
)

type pg struct {
	*sqlx.DB
	*pgx.Conn
	connConfig *pgx.ConnConfig
	logger     *zap.Logger
}

// NewPG opens a postgres db
func NewPG(uri string, logger *zap.Logger) (*pg, error) {
	cc, err := pgx.ParseConfig(uri)
	if err != nil {
		return nil, err
	}

	if cc.Port == 0 {
		cc.Port = defaultPgPort
	}

	db, err := sqlx.Open("pgx", uri)
	if err != nil {
		return nil, err
	}

	return &pg{
		DB:         db,
		connConfig: cc,
		logger:     logger,
	}, nil
}

// Connect connects to the database and verify with a ping, if successful - create db if not exist
func (p *pg) Connect(ctx context.Context) error {
	var err error
	var SSLmode string

	if p.connConfig.TLSConfig == nil {
		SSLmode = "disable"
	}

	p.DB, err = sqlx.ConnectContext(
		ctx,
		"pgx",
		fmt.Sprintf(
			connStr,
			p.connConfig.Host,
			p.connConfig.Port,
			p.connConfig.User,
			p.connConfig.Password,
			SSLmode,
		),
	)
	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	_, err = p.DB.Exec(fmt.Sprintf("CREATE DATABASE %s", p.connConfig.Database))
	if err != nil {
		var pgErr *pgconn.PgError
		if !errors.As(err, &pgErr) || pgerrcode.DuplicateDatabase != pgErr.Code {
			p.logger.Error(err.Error())
			return err
		}
		p.logger.Info(err.Error())
	} else {
		p.logger.Info(fmt.Sprintf(logDatabaseCreated, p.connConfig.Database))
	}

	p.DB, err = sqlx.ConnectContext(
		ctx,
		"pgx",
		fmt.Sprintf(
			connStrDB,
			p.connConfig.Host,
			p.connConfig.Port,
			p.connConfig.User,
			p.connConfig.Password,
			p.connConfig.Database,
			SSLmode,
		),
	)
	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	p.Conn, err = pgx.ConnectConfig(ctx, p.connConfig)
	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	return nil
}

// Migrate applies all up migrations
func (p *pg) Migrate() error {
	d, err := postgres.WithInstance(p.DB.DB, &postgres.Config{})
	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://schema",
		p.connConfig.Database, d)
	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		p.logger.Error(err.Error())
		return err
	}

	return nil
}

func (p *pg) Disconnect(ctx context.Context) error {
	return p.DB.Close()
}

func (p *pg) Ping(ctx context.Context) error {
	return p.PingContext(ctx)
}

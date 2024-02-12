package run

import (
	"context"
	"fmt"
	"net"
	"time"

	config "github.com/Chystik/pass-man/config/server"
	vaultcrypto "github.com/Chystik/pass-man/internal/crypto"
	"github.com/Chystik/pass-man/internal/infrastructure/db"
	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	"github.com/Chystik/pass-man/internal/interceptors"
	adapters "github.com/Chystik/pass-man/internal/user/adapters/server"
	"github.com/Chystik/pass-man/internal/user/infrastructure/repository"
	"github.com/Chystik/pass-man/internal/user/usecases"
	cardAdapters "github.com/Chystik/pass-man/internal/vault/card/adapters/server"
	cardRepo "github.com/Chystik/pass-man/internal/vault/card/infrastructure/repository"
	cardUsecases "github.com/Chystik/pass-man/internal/vault/card/usecases"
	fileAdapters "github.com/Chystik/pass-man/internal/vault/file/adapters/server"
	fileRepo "github.com/Chystik/pass-man/internal/vault/file/infrastructure/repository"
	fileUsecases "github.com/Chystik/pass-man/internal/vault/file/usecases"
	noteAdapters "github.com/Chystik/pass-man/internal/vault/note/adapters/server"
	noteRepo "github.com/Chystik/pass-man/internal/vault/note/infrastructure/repository"
	noteUsecases "github.com/Chystik/pass-man/internal/vault/note/usecases"
	passAdapters "github.com/Chystik/pass-man/internal/vault/password/adapters/server"
	passRepo "github.com/Chystik/pass-man/internal/vault/password/infrastructure/repository"
	passUsecases "github.com/Chystik/pass-man/internal/vault/password/usecases"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

const (
	logGRPCServerStart            = "gRPC server started on port: %s"
	logGRPCServerStop             = "Stopped serving new gRPC connections"
	logSignalInterrupt            = "Interrupt signal. Shutdown"
	logGracefulGRPCServerShutdown = "Graceful shutdown of gRPC Server complete."
	logDBDisconnect               = "Graceful close connection for DB client complete."
)

const (
	pingTimeout     = 3 * time.Second
	shutdownTimeout = 5 * time.Second
)

func Server(ctx context.Context, cfg *config.ServerConfig) {
	// Init logger
	log, err := initLogger(cfg.LogLevel, cfg.LogOutPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Connect DB and perform all UP migrations
	pg, err := db.NewPG(cfg.DBuri, log)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctxPing, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	err = pg.Connect(ctxPing)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = pg.Migrate()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Init User key store
	keyStore := vaultcrypto.NewKeyStore()

	// Init Vault cryptor
	cryptor := vaultcrypto.NewVaultCryptor(keyStore)

	// Init repositories
	useRepository := repository.NewUserRepository(pg.DB, log)
	passwordRepository := passRepo.NewPasswordRepository(pg.DB, log, cryptor)
	cardRepository := cardRepo.NewCardRepository(pg.DB, log, cryptor)
	fileRepository := fileRepo.NewFileRepository(pg.DB, pg.Conn, log, cryptor)
	noteRepository := noteRepo.NewNoteRepository(pg.DB, log, cryptor)

	// Create usecases
	userUsecases := usecases.NewUserUsecases(useRepository, keyStore)
	passwordUsecases := passUsecases.NewPasswordUsecases(passwordRepository)
	cardUsecases := cardUsecases.NewCardUsecases(cardRepository)
	fileUsecases := fileUsecases.NewFileUsecases(fileRepository)
	noteUsecases := noteUsecases.NewNoteRepository(noteRepository)

	// Init gRPC server
	lis, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		log.Fatal(err.Error())
	}

	gs := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.UnaryServerLogger(log),
			interceptors.UnaryServerAuth(cfg.AuthSecretKey),
			interceptors.UnaryServerRecoverer(log),
		),
		grpc.ChainStreamInterceptor(
			interceptors.StreamServerLogger(log),
			interceptors.StreamServerAuth(cfg.AuthSecretKey),
			interceptors.StreamServerRecoverer(log),
		),
	)

	// Register gRPC methods
	pb.RegisterUserServiceServer(gs, adapters.NewUserHandlers(userUsecases, cfg.AuthSecretKey))
	pb.RegisterPasswordServiceServer(gs, passAdapters.NewPasswordHandlers(passwordUsecases))
	pb.RegisterCardServiceServer(gs, cardAdapters.NewCardHandlers(cardUsecases))
	pb.RegisterFileServiceServer(gs, fileAdapters.NewFileHandlers(fileUsecases))
	pb.RegisterNoteServiceServer(gs, noteAdapters.NewPasswordHandlers(noteUsecases))

	// Run gRPC server
	go func() {
		log.Info(fmt.Sprintf(logGRPCServerStart, cfg.Address))
		if err = gs.Serve(lis); err != nil {
			log.Fatal(err.Error())
		}
		log.Info(logGRPCServerStop)
	}()

	// Wait interrupt signal
	<-ctx.Done()

	log.Info(logSignalInterrupt)
	ctxShutdown, shutdown := context.WithTimeout(context.Background(), shutdownTimeout)
	defer shutdown()

	// Graceful shutdown gRPC server
	gs.GracefulStop()
	log.Info(logGracefulGRPCServerShutdown)

	// Gracefil disconnect DB client
	if err := pg.Disconnect(ctxShutdown); err != nil {
		log.Fatal(err.Error())
	}
	log.Info(logDBDisconnect)
}

func initLogger(level string, outPath ...string) (*zap.Logger, error) {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	cfg := zap.NewProductionConfig()
	cfg.Level = lvl
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	if outPath != nil {
		cfg.OutputPaths = append(outPath, "stderr")
	}

	zl, err := cfg.Build()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return zl, nil
}

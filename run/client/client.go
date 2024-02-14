package run

import (
	"context"
	"fmt"
	"time"

	config "github.com/Chystik/pass-man/config/client"
	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	useradapters "github.com/Chystik/pass-man/internal/user/adapters/client"
	"github.com/Chystik/pass-man/internal/user/entities"
	cardAdapters "github.com/Chystik/pass-man/internal/vault/card/adapters/client"
	fileAdapters "github.com/Chystik/pass-man/internal/vault/file/adapters/client"
	noteAdapters "github.com/Chystik/pass-man/internal/vault/note/adapters/client"
	passAdapters "github.com/Chystik/pass-man/internal/vault/password/adapters/client"
	"github.com/Chystik/pass-man/run/client/cli"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type client struct {
	login     string
	authToken entities.JWTtoken
}

func Client(ctx context.Context, cfg *config.ClientConfig) {
	c := &client{
		login: cfg.GetLogin(),
	}

	conn, err := grpc.Dial(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// User service
	uc := pb.NewUserServiceClient(conn)

	userAPI := useradapters.NewUserAPIClient(uc)

	signCtx, cancelSign := context.WithTimeout(ctx, 10*time.Second)
	defer cancelSign()

	if cfg.SignType == config.Signup {
		c.authToken, err = userAPI.SignUp(signCtx, cfg.GetLogin(), cfg.GetPassword())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Successfully registered with login \"%s\" !\n", c.login)
	} else if cfg.SignType == config.Login {
		c.authToken, err = userAPI.Login(signCtx, cfg.GetLogin(), cfg.GetPassword())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Successfully logged in with login \"%s\" !\n", c.login)
	} else {
		panic("unknown sign type")
	}

	// vault services
	pc := pb.NewPasswordServiceClient(conn)
	cc := pb.NewCardServiceClient(conn)
	fc := pb.NewFileServiceClient(conn)
	nc := pb.NewNoteServiceClient(conn)

	passwordAPI := passAdapters.NewPasswordAPIClient(pc)
	cardAPI := cardAdapters.NewCardAPIClient(cc)
	fileAPI := fileAdapters.NewFileAPIClient(fc)
	noteAPI := noteAdapters.NewNoteAPIClient(nc)

	type vaultAPI struct {
		passAdapters.PasswordAPIClient
		cardAdapters.CardAPIClient
		fileAdapters.FileAPIClient
		noteAdapters.NoteAPIClient
	}

	va := vaultAPI{
		passwordAPI,
		cardAPI,
		fileAPI,
		noteAPI,
	}

	// Add auth token to all future requests
	ctxAuth := metadata.AppendToOutgoingContext(ctx, "token", string(c.authToken))

	cli := cli.NewCli(va)
	cli.Main(ctxAuth)
}

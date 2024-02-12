package run

import (
	"context"
	"fmt"
	"time"

	config "github.com/Chystik/pass-man/config/client"
	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	useradapters "github.com/Chystik/pass-man/internal/user/adapters/client"
	"github.com/Chystik/pass-man/internal/user/entities"
	vaultadapters "github.com/Chystik/pass-man/internal/vault/adapters/client"
	"github.com/Chystik/pass-man/internal/vault/usecases"
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

	uc := pb.NewUserServiceClient(conn)

	userAPI := useradapters.NewUserAPIClient(conn, uc)

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

	pc := pb.NewPasswordServiceClient(conn)
	cc := pb.NewCardServiceClient(conn)
	fc := pb.NewFileServiceClient(conn)

	passwordAPI := vaultadapters.NewPasswordAPIClient(conn, pc)
	cardApi := vaultadapters.NewCardAPIClient(conn, cc)
	fileApi := vaultadapters.NewFileAPIClient(conn, fc)

	type vaultAPI struct {
		usecases.PasswordAPIClient
		usecases.CardAPIClient
		usecases.FileAPIClient
	}

	va := vaultAPI{
		passwordAPI,
		cardApi,
		fileApi,
	}

	// Add auth token to all future requests
	//md := metadata.New(map[string]string{"token": string(c.authToken)})
	ctxAuth := metadata.AppendToOutgoingContext(ctx, "token", string(c.authToken))

	cli := cli.NewCli(va)
	cli.Main(ctxAuth)
}

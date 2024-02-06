package run

import (
	"context"
	"fmt"
	"time"

	config "github.com/Chystik/pass-man/config/client"
	pb "github.com/Chystik/pass-man/internal/infrastructure/grpc"
	adapters "github.com/Chystik/pass-man/internal/user/adapters/client"
	"github.com/Chystik/pass-man/internal/user/entities"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	authToken entities.JWTtoken
}

func Client(ctx context.Context, cfg *config.ClientConfig) {
	c := &client{}

	conn, err := grpc.Dial(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	uc := pb.NewUserServiceClient(conn)

	userAPI := adapters.NewUserAPIClient(conn, uc)
	signCtx, cancelSign := context.WithTimeout(ctx, 5*time.Second)
	defer cancelSign()

	if cfg.SignType == config.Signup {
		c.authToken, err = userAPI.SignUp(signCtx, cfg.GetLogin(), cfg.GetPassword())
		if err != nil {
			panic(err)
		}
	} else if cfg.SignType == config.Login {
		c.authToken, err = userAPI.Login(signCtx, cfg.GetLogin(), cfg.GetPassword())
		if err != nil {
			panic(err)
		}
	} else {
		panic("unknown sign type")
	}

	fmt.Println(c.authToken)
}

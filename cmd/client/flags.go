package main

import (
	"flag"
	"fmt"
	"os"

	config "github.com/Chystik/pass-man/config/client"
)

const (
	help = "expect -a=serverAddress `signup -u=login -p=password' or 'login -u=login -p=password'"
)

func parseFlags(cfg *config.ClientConfig) {
	addr := flag.NewFlagSet("Server address", flag.ExitOnError)
	addr.StringVar(&cfg.Address, "a", "", "server net address host:port")

	signUp := flag.NewFlagSet("Sign up", flag.ExitOnError)
	regUser := signUp.String("u", "", "User login")
	regPass := signUp.String("p", "", "User password")

	login := flag.NewFlagSet("Login", flag.ExitOnError)
	loginUser := login.String("u", "", "User login")
	loginPass := login.String("p", "", "User password")

	addr.Parse(os.Args[1:])
	if os.Args[1] != "-a" && os.Args[1] != "a" && os.Args[1] != "--a" {
		fmt.Fprintln(os.Stderr, help)
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, help)
		os.Exit(1)
	}

	switch os.Args[3] {
	case string(config.Signup):
		signUp.Parse(os.Args[4:])
		cfg.SetLogin(*regUser)
		cfg.SetPassword([]byte(*regPass))
		cfg.SignType = "signup"
	case string(config.Login):
		login.Parse(os.Args[4:])
		cfg.SetLogin(*loginUser)
		cfg.SetPassword([]byte(*loginPass))
		cfg.SignType = "login"
	default:
		fmt.Fprintln(os.Stderr, help)
		os.Exit(1)
	}
}

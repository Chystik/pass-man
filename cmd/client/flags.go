package main

import (
	"flag"
	"fmt"
	"os"

	config "github.com/Chystik/pass-man/config/client"
)

const (
	help = `Usage:

client [-a] "host:port" <command> [-u] "user" [-p] "password"
client [-v]
			
The commands are:

	signup		user registration
	login		user login

Use -v flag for print varsion and build information
			
Examples:

	Signup:	client -a=127.0.0.1:8080 signup -u=Elon -p=123456
	Login:	client -a=localhost:8080 login -u=Elon -p=123456`
)

var (
	buildVersion string
	buildDate    string
	buildCommit  string
)

func parseFlags(cfg *config.ClientConfig) {
	ver := flag.NewFlagSet("Build info", flag.ExitOnError)
	ver.Bool("v", false, "show binary version and build information")

	addr := flag.NewFlagSet("Server address", flag.ExitOnError)
	addr.StringVar(&cfg.Address, "a", "", "server net address host:port")

	signUp := flag.NewFlagSet("Sign up", flag.ExitOnError)
	regUser := signUp.String("u", "", "User login")
	regPass := signUp.String("p", "", "User password")

	login := flag.NewFlagSet("Login", flag.ExitOnError)
	loginUser := login.String("u", "", "User login")
	loginPass := login.String("p", "", "User password")

	if 2 > len(os.Args) || len(os.Args) > 8 {
		fmt.Fprintln(os.Stderr, help)
		os.Exit(1)
	}

	if os.Args[1] == "-v" || os.Args[1] == "v" || os.Args[1] == "--v" {
		fmt.Printf("Version: \t%s\nBuild date: \t%s\nBuild commit: \t%s\n",
			buildVersion, buildDate, buildCommit)
		os.Exit(0)
	}

	if len(os.Args) != 8 {
		fmt.Fprintln(os.Stderr, help)
		os.Exit(1)
	}

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

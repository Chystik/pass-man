package main

import (
	"os"

	config "github.com/Chystik/pass-man/config/server"

	"github.com/joho/godotenv"
)

func parseEnv(cfg *config.ServerConfig) {
	cfg.Address = os.Getenv("ADDRESS")
	cfg.LogLevel = os.Getenv("LOG_LEVEL")
	cfg.LogOutPath = os.Getenv("LOG_OUT_PATH")
	cfg.DBuri = os.Getenv("DATABASE_URI")
	cfg.AuthSecretKey = []byte(os.Getenv("AUTH_SECRET_KEY"))
}

func loadEnvFromFile(filePath string) error {
	return godotenv.Load(filePath)
}

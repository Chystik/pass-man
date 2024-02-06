package config

type ServerConfig struct {
	Address       string `env:"ADDRESS"`
	DBuri         string `env:"DATABASE_URI"`
	LogLevel      string `env:"LOG_LEVEL"`
	LogOutPath    string `env:"LOG_OUT_PATH"`
	AuthSecretKey []byte `env:"AUTH_SECRET_KEY"`
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Address:  ":8080",
		LogLevel: "info",
	}
}

package config

type signType string

const (
	Signup signType = "signup"
	Login  signType = "login"
)

type ClientConfig struct {
	SignType signType
	Address  string
	login    string
	password []byte
}

func NewClientConfig() *ClientConfig {
	return &ClientConfig{
		Address: ":8080",
	}
}

func (cc *ClientConfig) SetLogin(l string) {
	cc.login = l
}

func (cc *ClientConfig) SetPassword(p []byte) {
	cc.password = p
}

func (cc ClientConfig) GetLogin() string {
	return cc.login
}

func (cc ClientConfig) GetPassword() []byte {
	return cc.password
}

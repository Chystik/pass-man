package entities

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	ClaimsKey string
	JWTtoken  string
)

const (
	ClaimsKeyName   ClaimsKey = "props"
	TokenExpiration           = 24 * time.Hour
)

type AuthClaims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

func (ac *AuthClaims) AuthorizeUser(jwtKey []byte) (string, error) {
	expirationTime := time.Now().Add(TokenExpiration)

	ac.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, ac)

	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

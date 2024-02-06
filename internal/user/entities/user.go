package entities

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	appErr "github.com/Chystik/pass-man/internal/error/entities"

	"golang.org/x/crypto/bcrypt"
)

const (
	vaultKeyLen = 32
)

type User struct {
	Login             string `db:"login"`
	HashedPassword    []byte `db:"password"`
	EncryptedVaultKey []byte `db:"vault_key"`
}

// SetPassword hashes the user's password
func (u *User) SetPassword(password []byte) error {
	hash, err := bcrypt.GenerateFromPassword(password, 8)
	if err != nil {
		return err
	}
	u.HashedPassword = hash
	return nil
}

func (u User) Authenticate(password []byte) error {
	err := bcrypt.CompareHashAndPassword(u.HashedPassword, password)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return &appErr.AppError{Op: "user.Authenticate", Code: appErr.ErrUserCreds, Message: err.Error()}
		}
		return err
	}

	return nil
}

func (u User) GetLoginFromContext(ctx context.Context) (string, error) {
	claims, ok := ctx.Value(ClaimsKeyName).(*AuthClaims)
	if !ok {
		return "", &appErr.AppError{Op: "user.GetLoginFromContext", Code: appErr.ErrAuthClaims}
	}

	return claims.Login, nil
}

func (u *User) SetVaultKey(vaultPassword []byte) error {
	if len(vaultPassword) > vaultKeyLen {
		return &appErr.AppError{Op: "user.SetVaultKey", Code: appErr.ErrAuthClaims, Message: "vault password > 256 bit"}
	}

	// Generate a random 32 byte key for AES-256
	vaultKey := make([]byte, vaultKeyLen)

	_, err := rand.Read(vaultKey)
	if err != nil {
		return err
	}

	key := make([]byte, 32)
	copy(key[:len(vaultPassword)], vaultPassword)

	// Generate a new aes cipher using our 32 byte long key
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())

	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	// Encrypt the data using aesGCM.Seal
	// Since we don't want to save the nonce somewhere else in this case,
	// we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	cipherKey := gcm.Seal(nonce, nonce, vaultKey, nil)

	u.EncryptedVaultKey = cipherKey
	return nil
}

func (u *User) GetVaultKey(vaultPassword []byte) ([]byte, error) {
	if len(vaultPassword) > vaultKeyLen {
		return nil, &appErr.AppError{Op: "user.SetVaultKey", Code: appErr.ErrAuthClaims, Message: "vault password > 256 bit"}
	}

	key := make([]byte, 32)
	copy(key[:len(vaultPassword)], vaultPassword)

	// Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Create a new GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Get the nonce size
	nonceSize := gcm.NonceSize()

	// Extract the nonce from the encrypted data
	nonce, cipherKey := u.EncryptedVaultKey[:nonceSize], u.EncryptedVaultKey[nonceSize:]

	// Decrypt the data
	vaultKey, err := gcm.Open(nil, nonce, cipherKey, nil)
	if err != nil {
		return nil, err
	}

	return vaultKey, nil
}

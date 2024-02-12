package vault

import "io"

type VaultCryptor interface {
	Encrypt(in io.Reader, out io.Writer, userID string) (int, error)
	Decrypt(in io.Reader, out io.Writer, userID string) (int, error)
}

type VaultKeyStore interface {
	Lock(login string) error
	Unlock(login string, key []byte) error
	GetKey(login string) ([]byte, error)
}

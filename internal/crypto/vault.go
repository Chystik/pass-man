package vaultcrypto

import (
	"io"

	"github.com/Chystik/pass-man/internal/vault/usecases"
)

type vaultCryptor struct {
	key usecases.VaultKeyStore
}

func NewVaultCryptor(keyStore usecases.VaultKeyStore) *vaultCryptor {
	return &vaultCryptor{
		key: keyStore,
	}
}

type cryptoReader struct {
	r io.Reader
}

type cryptoWriter struct {
	w io.Writer
}

func newDecryptor() {
	panic("TODO")
}

func newEncryptor() {
	panic("TODO")
}

func (cr *cryptoReader) Read(p []byte) (n int, err error) {
	panic("TODO")
}

func (cw *cryptoWriter) Write(p []byte) (n int, err error) {
	panic("TODO")
}

func (v *vaultCryptor) Encrypt(p []byte) (n int, err error) {
	panic("TODO")
}

func (v *vaultCryptor) Decrypt(p []byte) (n int, err error) {
	panic("TODO")
}

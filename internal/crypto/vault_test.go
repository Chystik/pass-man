package vaultcrypto

import (
	"bytes"
	"testing"

	"github.com/Chystik/pass-man/internal/user/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVaultCryptor(t *testing.T) {
	t.Parallel()

	testData := "this is text for testing encryptor"

	keyStore := NewKeyStore()
	user := entities.User{
		Login: "test",
	}
	pass := []byte("userPassword")

	err := user.SetPassword(pass)
	require.NoError(t, err)

	err = user.SetVaultKey(pass)
	require.NoError(t, err)

	vaultKey, err := user.GetVaultKey(pass)
	require.NoError(t, err)

	err = keyStore.Unlock(user.Login, vaultKey)
	require.NoError(t, err)

	cryptor := NewVaultCryptor(keyStore)

	in := &bytes.Buffer{}
	in.Write([]byte(testData))

	out := &bytes.Buffer{}

	n, err := cryptor.Encrypt(in, out, user.Login)
	require.NoError(t, err)

	encrypted := out.Bytes()

	assert.Equal(t, len(encrypted)-1, n)

	inEncrypted := &bytes.Buffer{}
	inEncrypted.Write(encrypted)

	outDecrypted := &bytes.Buffer{}

	r, err := cryptor.Decrypt(inEncrypted, outDecrypted, user.Login)
	require.NoError(t, err)

	decrypted := outDecrypted.String()

	assert.Equal(t, len(decrypted)+16, r)
	assert.Equal(t, testData, decrypted)
}

package vaultcrypto

import (
	"sync"

	"github.com/Chystik/pass-man/internal/error/entities"
)

type keyStore struct {
	mu    sync.RWMutex
	store map[string][]byte
}

func NewKeyStore() *keyStore {
	return &keyStore{
		store: make(map[string][]byte),
	}
}

func (v *keyStore) Lock(login string) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	if v, ok := v.store[login]; ok {
		for i := range v {
			v[i] = 0
		}
	}

	return nil
}

func (v *keyStore) Unlock(login string, key []byte) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	v.store[login] = key

	return nil
}

func (v *keyStore) GetKey(login string) ([]byte, error) {
	v.mu.RLock()
	defer v.mu.RUnlock()

	k, ok := v.store[login]
	if !ok {
		return nil, &entities.AppError{Op: "vaultcrypto.GetKey", Code: entities.ErrNotFound}
	}

	return k, nil
}

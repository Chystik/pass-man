package repository

import (
	"bytes"

	"github.com/Chystik/pass-man/internal/vault/entities"
)

type dsFile struct {
	ID     uint32 `db:"id"`
	UserID string `db:"user_id"`
	Meta   string `db:"meta"`
	Name   []byte `db:"full_name"`
}

func (fr *fileRepository) fromDomainFile(f *entities.File, userID string) (*dsFile, error) {
	res := &dsFile{
		ID:     f.ID,
		UserID: userID,
		Meta:   f.Meta,
	}
	var err error
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}

	_, err = in.Write([]byte(f.Name))
	if err != nil {
		return res, err
	}

	_, err = fr.cryptor.Encrypt(in, out, userID)
	if err != nil {
		return res, err
	}

	res.Name = out.Bytes()

	return res, nil
}

func (fr *fileRepository) toDomainFile(f *dsFile, userID string) (*entities.File, error) {
	res := &entities.File{
		ID:   f.ID,
		Meta: f.Meta,
	}
	var err error
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}

	_, err = in.Write(f.Name)
	if err != nil {
		return res, err
	}

	_, err = fr.cryptor.Decrypt(in, out, userID)
	if err != nil {
		return res, err
	}

	res.Name = out.String()

	return res, nil
}

func (fr *fileRepository) toDomainFiles(f []*dsFile, userID string) ([]*entities.File, error) {
	var err error
	res := make([]*entities.File, len(f))

	for i := range f {
		res[i], err = fr.toDomainFile(f[i], userID)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

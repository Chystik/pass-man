package repository

import (
	"bytes"

	"github.com/Chystik/pass-man/internal/vault/entities"
)

type dsPassword struct {
	UserID   string `db:"user_id"`
	Meta     string `db:"meta"`
	Username []byte `db:"username"`
	Password []byte `db:"password"`
}

func (r *passwordRepository) fromDomainPassword(p entities.Password, userID string) (dsPassword, error) {
	var err error
	res := dsPassword{
		UserID: userID,
		Meta:   p.Meta,
	}

	res.Username, err = r.encrypt(p.Username, userID)
	if err != nil {
		return res, err
	}

	res.Password, err = r.encrypt(p.Password, userID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *passwordRepository) toDomainPassword(p dsPassword, userID string) (entities.Password, error) {
	var err error
	res := entities.Password{
		Meta: string(p.Meta),
	}

	res.Username, err = r.decrypt(p.Username, userID)
	if err != nil {
		return res, err
	}

	res.Password, err = r.decrypt(p.Password, userID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *passwordRepository) toDomainPasswords(p []dsPassword, userID string) ([]entities.Password, error) {
	var err error
	res := make([]entities.Password, len(p))

	for i := range p {
		res[i], err = r.toDomainPassword(p[i], userID)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

func (r *passwordRepository) encrypt(d string, id string) ([]byte, error) {
	var err error
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}

	_, err = in.Write([]byte(d))
	if err != nil {
		return nil, err
	}

	_, err = r.cryptor.Encrypt(in, out, id)
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func (r *passwordRepository) decrypt(d []byte, id string) (string, error) {
	var err error
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}

	_, err = in.Write(d)
	if err != nil {
		return "", err
	}

	_, err = r.cryptor.Decrypt(in, out, id)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

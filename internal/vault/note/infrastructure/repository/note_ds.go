package repository

import (
	"bytes"

	"github.com/Chystik/pass-man/internal/vault/note/entities"
)

type dsNote struct {
	UserID string `db:"user_id"`
	Meta   string `db:"meta"`
	Note   []byte `db:"note"`
}

func (r *noteRepository) fromDomainNote(p entities.Note, userID string) (dsNote, error) {
	var err error
	res := dsNote{
		UserID: userID,
		Meta:   p.Meta,
	}

	res.Note, err = r.encrypt(p.Note, userID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *noteRepository) toDomainNote(p dsNote, userID string) (entities.Note, error) {
	var err error
	res := entities.Note{
		Meta: string(p.Meta),
	}

	res.Note, err = r.decrypt(p.Note, userID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *noteRepository) toDomainNotes(p []dsNote, userID string) ([]entities.Note, error) {
	var err error
	res := make([]entities.Note, len(p))

	for i := range p {
		res[i], err = r.toDomainNote(p[i], userID)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

func (r *noteRepository) encrypt(d string, id string) ([]byte, error) {
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

func (r *noteRepository) decrypt(d []byte, id string) (string, error) {
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

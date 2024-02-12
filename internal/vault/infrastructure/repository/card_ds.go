package repository

import (
	"bytes"
	"strconv"

	"github.com/Chystik/pass-man/internal/vault/entities"
)

type dsCard struct {
	UserID    string `db:"user_id"`
	Meta      string `db:"meta"`
	Number    []byte `db:"number"`
	ValidThru []byte `db:"valid_thru"`
	Holder    []byte `db:"holder"`
	CVV       []byte `db:"cvv"`
}

func (r *cardRepository) fromDomainCard(c entities.Card, userID string) (dsCard, error) {
	var err error
	res := dsCard{
		UserID: userID,
		Meta:   c.Meta,
	}

	res.Number, err = r.encrypt(c.Number, userID)
	if err != nil {
		return res, err
	}

	res.ValidThru, err = r.encrypt(c.ValidThru.String(), userID)
	if err != nil {
		return res, err
	}

	res.Holder, err = r.encrypt(c.Holder, userID)
	if err != nil {
		return res, err
	}

	res.CVV, err = r.encrypt(strconv.Itoa(c.CVV), userID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *cardRepository) toDomainCard(c dsCard, userID string) (entities.Card, error) {
	var err error
	res := entities.Card{
		Meta: string(c.Meta),
	}

	res.Number, err = r.decrypt(c.Number, userID)
	if err != nil {
		return res, err
	}

	vt, err := r.decrypt(c.ValidThru, userID)
	if err != nil {
		return res, err
	}

	err = res.ValidThru.Parse(vt)
	if err != nil {
		return res, err
	}

	res.Holder, err = r.decrypt(c.Holder, userID)
	if err != nil {
		return res, err
	}

	cvv, err := r.decrypt(c.CVV, userID)
	if err != nil {
		return res, err
	}

	res.CVV, err = strconv.Atoi(cvv)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *cardRepository) toDomainCards(c []dsCard, userID string) ([]entities.Card, error) {
	var err error
	res := make([]entities.Card, len(c))

	for i := range c {
		res[i], err = r.toDomainCard(c[i], userID)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

func (r *cardRepository) encrypt(d string, id string) ([]byte, error) {
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

func (r *cardRepository) decrypt(d []byte, id string) (string, error) {
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

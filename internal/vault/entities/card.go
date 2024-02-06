package entities

import "time"

type Card struct {
	UserID    string    `db:"user_id"`
	Meta      string    `db:"meta"`
	Number    string    `db:"number"`
	ValidThru time.Time `db:"valid_thru"`
	Holder    string    `db:"holder"`
	CVV       int       `db:"cvv"`
}

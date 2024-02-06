package entities

type Password struct {
	Meta     string `db:"meta"`
	Username string `db:"username"`
	Password string `db:"password"`
}

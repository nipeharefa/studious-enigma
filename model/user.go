package model

type (
	User struct {
		ID       int    `db:"id"`
		Email    string `db:"email"`
		Address  string `db:"address"`
		Password string `db:"password"`
	}
)

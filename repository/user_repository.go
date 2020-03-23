package repository

import "github.com/nipeharefa/lemonilo/model"

type (
	UserRepository interface {
		Create(*model.User) error
		FindOneByEmail(string) (model.User, error)
		Update(*model.User) error
	}
)

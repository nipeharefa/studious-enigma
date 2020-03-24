package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nipeharefa/lemonilo/model"
)

type (
	UserRepository interface {
		Create(*model.User) error
		FindOne(ID int) (model.User, error)
		FindOneByEmail(string) (model.User, error)
		Update(*model.User) error
	}

	userRepository struct {
		db *sqlx.DB
	}
)

func NewUserRepository(db *sqlx.DB) UserRepository {

	return userRepository{db}
}

func (ur userRepository) FindOne(ID int) (model.User, error) {

	user := model.User{}
	sql := "select * FROM user_account where id=$1"

	err := ur.db.Get(&user, sql, ID)

	return user, err

}

func (ur userRepository) Create(u *model.User) error {

	sql := "insert into user_account(email, password, address) values($1, $2, $3) returning id"

	err := ur.db.QueryRow(sql, u.Email, u.Password, u.Address).Scan(&u.ID)

	return err
}

func (ur userRepository) FindOneByEmail(email string) (model.User, error) {

	u := model.User{}

	sql := "select * from user_account where email=$1 limit 1"

	err := ur.db.Get(&u, sql, email)
	return u, err
}

func (ur userRepository) Update(*model.User) error {

	return nil
}

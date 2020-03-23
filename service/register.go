package service

import (
	"database/sql"
	"errors"

	"github.com/nipeharefa/lemonilo/model"
	"github.com/nipeharefa/lemonilo/repository"
)

type (
	RegisterData interface {
		GetEmail() string
		GetPassword() string
		GetName() string
		GetAddress() string
	}

	RegisterService interface {
		Register(RegisterData) error
	}

	registerService struct {
		userRepo repository.UserRepository
	}
)

var (
	ErrAccountExist = errors.New("account already exist")
)

func NewRegisterService() RegisterService {

	r := registerService{}

	return r
}

func (r registerService) Register(data RegisterData) error {

	// Find exist email

	user, err := r.userRepo.FindOneByEmail(data.GetEmail())

	if err != nil || err != sql.ErrNoRows {
		return err
	}

	if user.ID != 0 {
		return ErrAccountExist
	}

	// create user
	user = model.User{}
	user.Name = data.GetName()
	user.Email = data.GetEmail()
	user.Address = data.GetAddress()
	user.Password = data.GetPassword()

	err = r.userRepo.Create(&user)
	if err != nil {
		return err
	}

	return nil
}

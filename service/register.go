package service

import (
	"database/sql"
	"errors"

	"github.com/nipeharefa/lemonilo/model"
	"github.com/nipeharefa/lemonilo/repository"
	"golang.org/x/crypto/bcrypt"
)

type (
	RegisterData interface {
		GetEmail() string
		GetPassword() string
		// GetName() string
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

func NewRegisterService(userRepo repository.UserRepository) RegisterService {

	r := registerService{userRepo}

	return r
}

func (r registerService) Register(data RegisterData) error {

	// Find exist email
	user, err := r.userRepo.FindOneByEmail(data.GetEmail())
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if user.ID != 0 {
		return ErrAccountExist
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.GetPassword()), 10)
	// create user
	user = model.User{}
	user.Email = data.GetEmail()
	user.Address = data.GetAddress()
	user.Password = string(hashedPassword)

	err = r.userRepo.Create(&user)
	if err != nil {
		return err
	}

	return nil
}

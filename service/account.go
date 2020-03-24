package service

import (
	"github.com/nipeharefa/lemonilo/model"
	"github.com/nipeharefa/lemonilo/repository"
)

type (
	AccountService interface {
		FindUserById(int) (model.User, error)
		Update(int, UpdateAccountData) (model.User, error)
	}

	UpdateAccountData interface {
		GetAddress() string
		GetEmail() string
	}

	accountService struct {
		userRepo repository.UserRepository
	}
)

func NewAccountService(userRepo repository.UserRepository) AccountService {
	return accountService{userRepo}
}

func (as accountService) FindUserById(ID int) (model.User, error) {

	return as.userRepo.FindOne(ID)
}

func (as accountService) Update(ID int, data UpdateAccountData) (model.User, error) {

	user, err := as.userRepo.FindOne(ID)
	if err != nil {
		return user, nil
	}

	user.Address = data.GetAddress()
	user.Email = data.GetEmail()

	err = as.userRepo.Update(&user)

	return user, err
}

package service

import (
	"github.com/nipeharefa/lemonilo/model"
	"github.com/nipeharefa/lemonilo/repository"
)

type (
	AccountService interface {
		FindUserById(int) (model.User, error)
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

package service

import "github.com/nipeharefa/lemonilo/repository"

type (
	LoginData interface {
		GetEmail() string
		GetPassword() string
	}
	LoginService interface {
		Login(LoginData) error
	}

	loginService struct {
		userRepo repository.UserRepository
	}
)

func NewLoginService() LoginService {

	l := loginService{}
	return l
}

func (l loginService) Login(data LoginData) error {

	// find user
	_, err := l.userRepo.FindOneByEmail(data.GetEmail())
	if err != nil {
		return err
	}

	return nil
}

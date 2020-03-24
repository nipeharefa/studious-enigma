package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nipeharefa/lemonilo/repository"
	"golang.org/x/crypto/bcrypt"
)

type (
	LoginData interface {
		GetEmail() string
		GetPassword() string
	}
	LoginService interface {
		Login(LoginData) (string, error)
	}

	loginService struct {
		userRepo repository.UserRepository
	}
)

var (
	ErrLoginFail = errors.New("login fail")
)

func NewLoginService(userRepo repository.UserRepository) LoginService {

	l := loginService{userRepo}
	return l
}

func (l loginService) Login(data LoginData) (string, error) {

	var accessToken string
	// find user
	user, err := l.userRepo.FindOneByEmail(data.GetEmail())
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.GetPassword()))
	if err != nil {
		return accessToken, ErrLoginFail
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	accessToken, err = token.SignedString([]byte("secret"))
	if err != nil {
		return accessToken, ErrLoginFail
	}

	return accessToken, nil
}

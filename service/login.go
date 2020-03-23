package service

type (
	LoginData interface {
		GetEmail() string
		GetPassword() string
	}
	LoginService interface {
		Login() error
	}
)

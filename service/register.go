package service

ype (
	RegisterData interface {
		GetEmail() string
		GetPassword() string
		GetName() string
		GetAddress() string
	}
	
	RegisterService interface {
		Register() error
	}
)

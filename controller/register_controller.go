package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nipeharefa/lemonilo/service"
)

type (
	RegisterController interface {
		Register(echo.Context) error
	}

	registerController struct {
		registerService service.RegisterService
	}

	registerRequest struct {
		Address  string `json:"address"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	registerResponse struct {
		Message string `json:"message"`
	}
)

func NewRegisterController(registerService service.RegisterService) RegisterController {

	r := registerController{registerService}

	return r
}

func (r registerController) Register(c echo.Context) error {

	request := registerRequest{}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	err := r.registerService.Register(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	resMessage := registerResponse{Message: "Register berhasil"}
	return c.JSON(http.StatusCreated, resMessage)
}

func (rr registerRequest) GetAddress() string {
	return rr.Address
}

func (rr registerRequest) GetEmail() string {
	return rr.Email
}

func (rr registerRequest) GetPassword() string {
	return rr.Password
}

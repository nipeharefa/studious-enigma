package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nipeharefa/lemonilo/service"
)

type (
	LoginController interface {
		Login(echo.Context) error
	}

	loginController struct {
		loginService service.LoginService
	}

	loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func NewLoginController() LoginController {

	lc := loginController{}
	return lc
}

func (lc loginController) Login(c echo.Context) error {

	req := loginRequest{}

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	err := lc.loginService.Login(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	return c.JSON(http.StatusOK, nil)
}

func (lr loginRequest) GetEmail() string {

	return lr.Email
}

func (lr loginRequest) GetPassword() string {

	return lr.Password
}

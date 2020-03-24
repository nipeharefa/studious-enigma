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

	loginResponse struct {
		AccessToken string `json:"accessToken"`
	}
)

func NewLoginController(loginService service.LoginService) LoginController {

	lc := loginController{loginService}
	return lc
}

func (lc loginController) Login(c echo.Context) error {

	req := loginRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	accessToken, err := lc.loginService.Login(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	return c.JSON(http.StatusOK, loginResponse{AccessToken: accessToken})
}

func (lr loginRequest) GetEmail() string {

	return lr.Email
}

func (lr loginRequest) GetPassword() string {

	return lr.Password
}

package controller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/nipeharefa/lemonilo/service"
)

type (
	GetProfileController interface {
		GetProfile(echo.Context) error
	}

	getProfileController struct {
		accountService service.AccountService
	}

	getProfileResponse struct {
		ID      int    `json:"id"`
		Address string `json:"address"`
		Email   string `json:"email"`
	}
)

func NewGetProfileController(accountService service.AccountService) GetProfileController {

	return getProfileController{accountService}
}

func (g getProfileController) GetProfile(c echo.Context) error {

	userID := getUserIDFromContext(c)

	user, _ := g.accountService.FindUserById(userID)

	resData := getProfileResponse{}
	resData.Address = user.Address
	resData.Email = user.Email
	resData.ID = user.ID

	return c.JSON(http.StatusOK, resData)
}

func getUserIDFromContext(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["user_id"]

	return int(id.(float64))
}

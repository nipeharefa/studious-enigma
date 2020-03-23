package controller

import "github.com/labstack/echo/v4"

type (
	GetProfileController interface {
		GetProfile(echo.Context) error
	}
)

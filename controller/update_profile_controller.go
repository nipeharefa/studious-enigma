package controller

import "github.com/labstack/echo/v4"

type (
	UpdateProfileController interface {
		Update(echo.Context) error
	}
)

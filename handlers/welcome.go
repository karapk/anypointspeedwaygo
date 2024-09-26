package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func WelcomeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Anypoint racing! 🚗💨" )
}

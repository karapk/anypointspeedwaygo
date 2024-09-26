package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Anypoint Racing! 🚗💨")
	})

	e.ServeHTTP(w, r)
}

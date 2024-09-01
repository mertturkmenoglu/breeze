package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ApiLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	fmt.Println(email, password)

	c.Response().Header().Set("HX-Redirect", "/")
	return c.Redirect(http.StatusSeeOther, "http://localhost:5000/")
}

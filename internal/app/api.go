package app

import (
	"breeze/internal/partials"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ApiLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	urls, _ := c.FormParams()

	fmt.Println(email, password, urls)

	c.Response().Header().Set("HX-Redirect", "/")
	return c.Redirect(http.StatusSeeOther, "http://localhost:5000/")
}

func ApiRegister(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	name := c.FormValue("name")

	errs := make([]string, 0)

	if !isValidEmail(email) {
		errs = append(errs, "Invalid email address.")
	}

	if !isValidPassword(password) {
		errs = append(errs, "Password must be at least 8 characters long.")
	}

	if !isValidName(name) {
		errs = append(errs, "Name must be at least 2 characters long.")
	}

	if len(errs) > 0 {
		return Render(c, http.StatusOK, partials.AuthErr(errs))
	}

	fmt.Println(email, password, name)

	c.Response().Header().Set("HX-Redirect", "/")
	return c.Redirect(http.StatusSeeOther, "http://localhost:5000/")
}

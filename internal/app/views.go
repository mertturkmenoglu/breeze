package app

import (
	"breeze/internal/db"
	"breeze/internal/views"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

// GET /
func HomeHandler(c echo.Context) error {
	d := db.NewDb()

	s, err := d.Queries.Dummy(c.Request().Context())

	if err != nil {
		return err
	}

	return Render(c, http.StatusOK, views.Home(s))
}

// GET /login
func LoginHandler(c echo.Context) error {
	return Render(c, http.StatusOK, views.Login())
}

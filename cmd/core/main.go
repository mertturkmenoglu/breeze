package main

import (
	"breeze/config"
	"breeze/internal/db"
	"breeze/internal/views"
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
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

func HomeHandler(c echo.Context) error {
	d := db.NewDb()

	s, err := d.Queries.Dummy(c.Request().Context())

	if err != nil {
		return err
	}

	return Render(c, http.StatusOK, views.Home(s))
}

func LoginHandler(c echo.Context) error {
	return Render(c, http.StatusOK, views.Login())
}

func main() {
	config.Bootstrap()

	shouldRunMigrations := os.Getenv("RUN_MIGRATIONS")

	if shouldRunMigrations == "1" {
		db.RunMigrations()
	}

	e := echo.New()

	e.Static("/assets", "internal/assets")

	e.GET("/", HomeHandler)
	e.GET("/login", LoginHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", viper.GetInt(config.PORT))))
}

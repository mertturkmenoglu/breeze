package main

import (
	"breeze/config"
	"breeze/internal/db"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	config.Bootstrap()

	shouldRunMigrations := os.Getenv("RUN_MIGRATIONS")

	if shouldRunMigrations == "1" {
		db.RunMigrations()
	}

	e := echo.New()

	d := db.NewDb()

	e.GET("/", func(c echo.Context) error {
		s, err := d.Queries.Dummy(c.Request().Context())

		if err != nil {
			return err
		}

		return c.String(200, s)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", viper.GetInt(config.PORT))))
}

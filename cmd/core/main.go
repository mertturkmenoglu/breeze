package main

import (
	"breeze/config"
	"breeze/internal/app"
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

	e.Static("/assets", "internal/assets")

	e.GET("/", app.HomeHandler)
	e.GET("/login", app.LoginHandler)
	e.POST("/login", app.ApiLogin)
	e.GET("/register", app.RegisterHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", viper.GetInt(config.PORT))))
}

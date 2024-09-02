package main

import (
	"breeze/config"
	"breeze/internal/app"
	"breeze/internal/db"
	"breeze/internal/middlewares"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	config.Bootstrap()

	shouldRunMigrations := os.Getenv("RUN_MIGRATIONS")

	if shouldRunMigrations == "1" {
		db.RunMigrations()
	}

	e := echo.New()

	e.Use(middlewares.PTermLogger)
	e.Use(middlewares.GetSessionMiddleware())

	e.Static("/assets", "internal/assets")

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLength:    32,
		TokenLookup:    "form:_csrf",
		CookieName:     "csrf_token",
		CookiePath:     "/",
		CookieHTTPOnly: true,
		CookieSecure:   true,
	}))

	db := db.NewDb()

	h := app.New(db)

	e.Use(middlewares.WithAuth)

	e.GET("/", h.HomeHandler)
	e.GET("/login", h.LoginHandler)
	e.POST("/login", h.LoginPostHandler)
	e.GET("/register", h.RegisterHandler)
	e.POST("/register", h.RegisterPostHandler)
	e.DELETE("/logout", h.LogoutHandler)
	e.GET("/new", h.NewHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", viper.GetInt(config.PORT))))
}

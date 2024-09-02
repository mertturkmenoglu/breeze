package app

import (
	"breeze/config"
	"breeze/internal/db"
	"breeze/internal/hash"
	"breeze/internal/partials"
	"breeze/internal/views"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/pterm/pterm"
	"github.com/spf13/viper"
)

type Handler struct {
	db *db.Db
}

func New(db *db.Db) *Handler {
	return &Handler{
		db: db,
	}
}

const SESSION_NAME = "__breeze_auth"

var logger = pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)

func getAuthSessionOptions() *sessions.Options {
	return &sessions.Options{
		Path:     viper.GetString(config.AUTH_SESSION_PATH),
		MaxAge:   viper.GetInt(config.AUTH_SESSION_MAX_AGE),
		HttpOnly: true,
		Secure:   viper.GetString(config.ENV) != "dev",
		SameSite: http.SameSiteLaxMode,
	}
}

// GET /
func (h *Handler) HomeHandler(c echo.Context) error {
	csrfToken := c.Get("csrf").(string)
	user, ok := c.Get("user").(db.User)
	var msg = ""

	if !ok || user.Name == "" {
		msg = "User"
	} else {
		msg = user.Name
	}

	isAuth := c.Get("user_id").(string) != ""

	return Render(c, http.StatusOK, views.Home(msg, isAuth, csrfToken))
}

// GET /login
func (h *Handler) LoginHandler(c echo.Context) error {
	csrfToken := c.Get("csrf").(string)
	return Render(c, http.StatusOK, views.Login(csrfToken))
}

// GET /register
func (h *Handler) RegisterHandler(c echo.Context) error {
	csrfToken := c.Get("csrf").(string)
	return Render(c, http.StatusOK, views.Register(csrfToken))
}

// POST /login
func (h *Handler) LoginPostHandler(c echo.Context) error {
	sess, err := session.Get(SESSION_NAME, c)

	if err != nil {
		logger.Error("error getting session", logger.Args("err", err))
		return echo.ErrInternalServerError
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	if !isValidEmail(email) || !isValidLoginPassword(password) {
		return Render(c, http.StatusOK, partials.AuthErr(
			[]string{
				"Invalid email address or password.",
			},
		))
	}

	user, err := h.db.Queries.GetUserByEmail(context.Background(), email)

	var hashed = ""

	if err == nil {
		hashed = user.PasswordHash
	}

	v, verifyErr := hash.Verify(password, hashed)

	if err != nil || verifyErr != nil || !v {
		return Render(c, http.StatusOK, partials.AuthErr(
			[]string{
				"Invalid email address or password.",
			},
		))
	}

	sessionId := uuid.New().String()
	createdAt := time.Now()
	expiresAt := createdAt.Add(time.Hour * 24 * 7)

	_, err = h.db.Queries.CreateSession(context.Background(), db.CreateSessionParams{
		ID:          sessionId,
		UserID:      user.ID,
		SessionData: pgtype.Text{},
		CreatedAt:   pgtype.Timestamptz{Time: createdAt, Valid: true},
		ExpiresAt:   pgtype.Timestamptz{Time: expiresAt, Valid: true},
	})

	if err != nil {
		return echo.ErrInternalServerError
	}

	sess.Options = getAuthSessionOptions()
	sess.Values["user_id"] = user.ID
	sess.Values["session_id"] = sessionId
	sess.Save(c.Request(), c.Response())

	return c.JSON(http.StatusOK, echo.Map{
		"redirect": "/",
	})
}

// POST /register
func (h *Handler) RegisterPostHandler(c echo.Context) error {
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

	_, err := h.db.Queries.GetUserByEmail(context.Background(), email)

	emailTaken := !(err != nil && errors.Is(err, pgx.ErrNoRows))

	if emailTaken {
		return Render(c, http.StatusOK, partials.AuthErr([]string{"Email is taken."}))
	}

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	hashed, err := hash.Hash(password)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	_, err = h.db.Queries.CreateUser(context.Background(), db.CreateUserParams{
		ID:           uuid.New().String(),
		Email:        email,
		PasswordHash: hashed,
		Name:         name,
		Role:         "user",
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"redirect": "/login",
	})
}

// DELETE /logout
func (h *Handler) LogoutHandler(c echo.Context) error {
	sess, err := session.Get(SESSION_NAME, c)

	if err != nil {
		return echo.ErrInternalServerError
	}

	delete(sess.Values, "user_id")
	sess.Save(c.Request(), c.Response())
	cookie := resetCookie()
	c.SetCookie(cookie)

	return c.NoContent(http.StatusNoContent)
}

func resetCookie() *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = SESSION_NAME
	cookie.Value = ""
	cookie.Path = "/"
	cookie.Expires = time.Unix(0, 0)
	cookie.MaxAge = -1

	return cookie
}

func (h *Handler) NewHandler(c echo.Context) error {
	return Render(c, http.StatusOK, views.New(c.Get("csrf").(string)))
}
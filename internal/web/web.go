package web

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Devil666face/fiber/internal/config"
	"github.com/Devil666face/fiber/internal/models"
	"github.com/Devil666face/fiber/internal/store/database"
	"github.com/Devil666face/fiber/internal/store/session"
	"github.com/Devil666face/fiber/internal/web/handlers"
	"github.com/Devil666face/fiber/internal/web/routes"
	"github.com/Devil666face/fiber/internal/web/validators"

	"github.com/gofiber/fiber/v2"
)

type Web struct {
	fiber     *fiber.App
	static    func(*fiber.Ctx) error
	media     *Media
	config    *config.Config
	database  *database.Database
	router    *routes.Router
	store     *session.Store
	validator *validators.Validator
	tables    []any
}

func New() *Web {
	a := &Web{
		fiber: fiber.New(
			fiber.Config{
				AppName:      "fiber",
				ErrorHandler: handlers.DefaultErrorHandler,
			},
		),
		static:    NewStatic(),
		media:     MustMedia(),
		config:    config.Must(),
		validator: validators.New(),
		tables: []any{
			&models.User{},
		},
	}
	a.setStores()
	a.setStatic()
	a.setRoutes()
	return a
}

func (a *Web) setStores() {
	a.database = database.Must(a.config, a.tables)
	a.store = session.New(a.config, a.database)
}

func (a *Web) setStatic() {
	a.fiber.Use(routes.StaticPrefix, a.static)
	a.fiber.Static(routes.MediaPrefix, a.media.path, a.media.handler)
}

func (a *Web) setRoutes() {
	a.router = routes.New(a.fiber, a.config, a.database, a.store, a.validator)
}

func (a *Web) Listen() error {
	if a.config.UseTLS {
		return a.listenTLS()
	}
	return a.listenNoTLS()
}

func (a *Web) listenTLS() error {
	go a.mustRedirectServer()
	return a.fiber.ListenTLS(a.config.ConnectHTTPS, a.config.TLSCrt, a.config.TLSKey)
}

func (a *Web) listenNoTLS() error {
	return a.fiber.Listen(a.config.ConnectHTTP)
}

func (a *Web) mustRedirectServer() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(func(c *fiber.Ctx) error {
		return c.Redirect(a.config.HTTPSRedirect)
	})
	if err := app.Listen(a.config.ConnectHTTP); err != nil {
		slog.Error(fmt.Sprintf("Start redirect server: %s", err))
		//nolint:revive //If connection for redirect server already busy - close app
		os.Exit(1)
	}
}

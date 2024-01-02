package routes

import (
	"github.com/Devil666face/fiber/internal/config"
	"github.com/Devil666face/fiber/internal/store/database"
	"github.com/Devil666face/fiber/internal/store/session"
	"github.com/Devil666face/fiber/internal/web/handlers"
	"github.com/Devil666face/fiber/internal/web/middlewares"
	"github.com/Devil666face/fiber/internal/web/validators"

	"github.com/gofiber/fiber/v2"
)

var (
	StaticPrefix = "/static"
	MediaPrefix  = "/media"
)

type Router struct {
	router      fiber.Router
	config      *config.Config
	database    *database.Database
	store       *session.Store
	validator   *validators.Validator
	middlewares []func(*handlers.Handler) error
}

func New(
	_router fiber.Router,
	_config *config.Config,
	_database *database.Database,
	_store *session.Store,
	_validator *validators.Validator,
) *Router {
	r := Router{
		router:    _router,
		config:    _config,
		database:  _database,
		store:     _store,
		validator: _validator,
		middlewares: []func(*handlers.Handler) error{
			middlewares.Logger,
			middlewares.Recover,
			middlewares.Compress,
			middlewares.Limiter,
			middlewares.AllowHost,
			middlewares.SecureHeaders,
			middlewares.EncryptCookie,
			middlewares.Csrf,
			middlewares.Htmx,
		},
	}
	r.setMiddlewares()
	r.setAuth()
	r.setUser()
	r.setIndex()
	return &r
}

func (r *Router) wrapper(handler func(*handlers.Handler) error) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return handler(handlers.New(c, r.database, r.config, r.store, r.validator))
	}
}

func (r *Router) setMiddlewares() {
	for _, middleware := range r.middlewares {
		r.router.Use(r.wrapper(middleware))
	}
}

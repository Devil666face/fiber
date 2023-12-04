package handlers

import (
	"github.com/Devil666face/fiber/internal/config"
	"github.com/Devil666face/fiber/internal/store/database"
	"github.com/Devil666face/fiber/internal/store/session"
	"github.com/Devil666face/fiber/internal/web/validators"
	"github.com/Devil666face/fiber/internal/web/view"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	fibersession "github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type Handler struct {
	c         *fiber.Ctx
	view      *view.View
	database  *database.Database
	config    *config.Config
	store     *session.Store
	validator *validators.Validator
	session   *fibersession.Session
}

func New(
	_c *fiber.Ctx,
	_database *database.Database,
	_config *config.Config,
	_store *session.Store,
	_validator *validators.Validator,
) *Handler {
	return &Handler{
		c:         _c,
		view:      view.New(_c),
		database:  _database,
		config:    _config,
		store:     _store,
		validator: _validator,
	}
}

func (h *Handler) Render(component func(*view.View, view.Map) templ.Component, m view.Map) error {
	h.c.Response().Header.SetContentType(fiber.MIMETextHTMLCharsetUTF8)
	return component(h.view, m).Render(h.c.UserContext(), h.c.Response().BodyWriter())
}

func (h *Handler) Ctx() *fiber.Ctx {
	return h.c
}

func (h *Handler) View() *view.View {
	return h.view
}

func (h *Handler) Database() *gorm.DB {
	return h.database.DB()
}

func (h *Handler) Store() *fibersession.Store {
	return h.store.Store()
}

func (h *Handler) Storage() fiber.Storage {
	return h.store.Storage()
}

func (h *Handler) Config() *config.Config {
	return h.config
}

func (h *Handler) Validator() *validators.Validator {
	return h.validator
}

func (h *Handler) getSession() error {
	var err error
	if h.session, err = h.Store().Get(h.c); err != nil {
		return err
	}
	return nil
}

func (h *Handler) SetInSession(key string, val any) error {
	if err := h.getSession(); err != nil {
		return err
	}
	h.session.Set(key, val)
	return h.SaveSession()
}

func (h *Handler) GetFromSession(key string) (any, error) {
	if err := h.getSession(); err != nil {
		return nil, err
	}
	return h.session.Get(key), nil
}

func (h *Handler) SaveSession() error {
	return h.session.Save()
}

func (h *Handler) DestroySession() error {
	return h.session.Destroy()
}

func (h *Handler) DestroySessionByID(sessID string) error {
	if sessID == "" {
		return nil
	}
	return h.Store().Delete(sessID)
}

func (h *Handler) SessionID() (string, error) {
	if err := h.getSession(); err != nil {
		return "", err
	}
	return h.session.ID(), nil
}

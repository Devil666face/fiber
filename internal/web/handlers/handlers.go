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
	viewctx    *ViewCtx
	database   *database.Database
	config     *config.Config
	session    *session.Store
	validator  *validators.Validator
	ctxsession *fibersession.Session
}

func New(
	c *fiber.Ctx,
	_database *database.Database,
	_config *config.Config,
	_session *session.Store,
	_validator *validators.Validator,
) *Handler {
	return &Handler{
		viewctx:   NewViewCtx(c),
		database:  _database,
		config:    _config,
		session:   _session,
		validator: _validator,
	}
}

// func (h *Handler) Render(component templ.Component) error {
// 	h.Ctx().Response().Header.SetContentType(fiber.MIMETextHTMLCharsetUTF8)
// 	return component.Render(h.Ctx().UserContext(), h.Ctx().Response().BodyWriter())
// }

func (h *Handler) RenderTempl(component func(*view.View, fiber.Map) templ.Component, m fiber.Map) error {
	h.Ctx().Response().Header.SetContentType(fiber.MIMETextHTMLCharsetUTF8)
	return component(view.New(h.Ctx()), m).Render(h.Ctx().UserContext(), h.Ctx().Response().BodyWriter())
}

func (h *Handler) ViewCtx() *ViewCtx {
	return h.viewctx
}

func (h *Handler) Ctx() *fiber.Ctx {
	return h.viewctx.Ctx
}

func (h *Handler) Database() *gorm.DB {
	return h.database.DB()
}

func (h *Handler) Store() *fibersession.Store {
	return h.session.Store()
}

func (h *Handler) Storage() fiber.Storage {
	return h.session.Storage()
}

func (h *Handler) Config() *config.Config {
	return h.config
}

func (h *Handler) Validator() *validators.Validator {
	return h.validator
}

func (h *Handler) getSession() error {
	var err error
	if h.ctxsession, err = h.Store().Get(h.Ctx()); err != nil {
		return err
	}
	return nil
}

func (h *Handler) SetInSession(key string, val any) error {
	if err := h.getSession(); err != nil {
		return err
	}
	h.ctxsession.Set(key, val)
	return h.SaveSession()
}

func (h *Handler) GetFromSession(key string) (any, error) {
	if err := h.getSession(); err != nil {
		return nil, err
	}
	return h.ctxsession.Get(key), nil
}

func (h *Handler) SaveSession() error {
	return h.ctxsession.Save()
}

func (h *Handler) DestroySession() error {
	return h.ctxsession.Destroy()
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
	return h.ctxsession.ID(), nil
}

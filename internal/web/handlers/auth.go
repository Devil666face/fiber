package handlers

import (
	"github.com/Devil666face/fiber/internal/models"
	"github.com/Devil666face/fiber/internal/web/view"

	"github.com/gofiber/fiber/v2"
)

var ErrInSession = fiber.ErrInternalServerError

func LoginPage(h *Handler) error {
	return h.RenderTempl(view.Login, view.Map{})
}

func Login(h *Handler) error {
	var (
		u   = &models.User{}
		in  = &models.User{}
		err error
	)
	if err := h.ViewCtx().BodyParser(in); err != nil {
		return err
	}
	u.Email = in.Email
	if err := u.LoginValidate(h.Database(), h.Validator(), in.Password); err != nil {
		return h.ViewCtx().RenderWithCtx("login", fiber.Map{
			"Title":   "Login",
			"Message": err.Error(),
		}, "base")
	}
	if err := h.SetInSession(view.AuthKey, true); err != nil {
		return ErrInSession
	}
	if err := h.SetInSession(view.UserID, u.ID); err != nil {
		return ErrInSession
	}
	if u.SessionKey, err = h.SessionID(); err != nil {
		return ErrInSession
	}
	if err := u.Update(h.Database()); err != nil {
		return ErrInSession
	}
	return h.ViewCtx().ClientRedirect(h.ViewCtx().URL("index"))
}

func Logout(h *Handler) error {
	if err := h.DestroySession(); err != nil {
		return ErrInSession
	}
	return h.ViewCtx().RedirectToRoute("login", nil)
}

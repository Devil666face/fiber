package handlers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Devil666face/fiber/internal/models"
	"github.com/Devil666face/fiber/internal/web/validators"
	"github.com/Devil666face/fiber/internal/web/view"

	"github.com/gofiber/fiber/v2"
)

func UserListPage(h *Handler) error {
	// && h.ViewCtx().IsHtmxCurrentURL()
	if h.ViewCtx().IsHtmx() {
		return h.ViewCtx().RenderWithCtx("user_content", fiber.Map{
			"Users": models.GetAllUsers(h.Database()),
		})
	}
	return h.RenderTempl(view.UserList, fiber.Map{
		"Title": "List of users",
		"Users": models.GetAllUsers(h.Database()),
	})
	// return h.ViewCtx().RenderWithCtx("user_list", fiber.Map{
	// 	"Title": "List of users",
	// 	"Users": models.GetAllUsers(h.Database()),
	// }, "base")
}

func UserEditForm(h *Handler) error {
	u := models.User{}
	id, err := strconv.Atoi(h.ViewCtx().Params("id"))
	if err != nil {
		return fiber.ErrNotFound
	}
	u.ID = uint(id)
	if err := u.Get(h.Database()); err != nil {
		return fiber.ErrNotFound
	}
	u.Password = ""
	return h.ViewCtx().RenderWithCtx("user_edit", fiber.Map{
		UserKey: u,
	})
}

func UserCreateForm(h *Handler) error {
	return h.ViewCtx().RenderWithCtx("user_create", fiber.Map{})
}

func UserCreate(h *Handler) error {
	u := models.User{}
	if err := h.ViewCtx().BodyParser(&u); err != nil {
		return fiber.ErrBadRequest
	}
	if err := u.Validate(h.Validator()); err != nil {
		return h.ViewCtx().RenderWithCtx("user_create", fiber.Map{
			"Message": err.Error(),
		})
	}
	if err := u.Create(h.Database()); err != nil {
		return h.ViewCtx().RenderWithCtx("user_create", fiber.Map{
			"Message": err.Error(),
		})
	}
	return h.ViewCtx().RenderWithCtx("user_create", fiber.Map{
		"Success": fmt.Sprintf("User %s - created", u.Email),
	})
}

func UserEdit(h *Handler) error {
	var (
		u  = models.User{}
		in = models.User{}
	)
	if err := h.ViewCtx().BodyParser(&in); err != nil {
		return err
	}
	id, err := strconv.Atoi(h.ViewCtx().Params("id"))
	if err != nil {
		return fiber.ErrNotFound
	}
	in.ID = uint(id)
	u.ID = in.ID
	if err := u.Get(h.Database()); err != nil {
		return fiber.ErrNotFound
	}
	if err := in.Validate(h.Validator()); err != nil {
		if errors.Is(err, validators.ErrPasswordRequired) {
			in.Password, in.PasswordConfirm = u.Password, u.Password
		} else {
			return h.ViewCtx().RenderWithCtx("user_edit", fiber.Map{
				UserKey:   u,
				"Message": err,
			})
		}
	}
	u.Email, u.Admin, u.Password = in.Email, in.Admin, in.Password
	if err := u.Update(h.Database()); err != nil {
		return h.ViewCtx().RenderWithCtx("user_edit", fiber.Map{
			UserKey:   u,
			"Message": err,
		})
	}
	if err := h.DestroySessionByID(u.SessionKey); err != nil {
		return h.ViewCtx().RenderWithCtx("user_edit", fiber.Map{
			UserKey:   u,
			"Message": err,
			// "Message":   ErrInSession,
		})
	}
	return h.ViewCtx().RenderWithCtx("user_edit", fiber.Map{
		UserKey:   u,
		"Success": "Successful update user",
	})
}

func UserDelete(h *Handler) error {
	u := models.User{}
	if err := h.ViewCtx().BodyParser(&u); err != nil {
		return err
	}
	id, err := strconv.Atoi(h.ViewCtx().Params("id"))
	if err != nil {
		return fiber.ErrNotFound
	}
	u.ID = uint(id)
	if err := u.Get(h.Database()); err != nil {
		return fiber.ErrNotFound
	}
	if err := u.Delete(h.Database()); err != nil {
		return err
	}
	if err := h.DestroySessionByID(u.SessionKey); err != nil {
		return ErrInSession
	}
	return h.ViewCtx().RenderWithCtx("user_content", fiber.Map{
		"Users": models.GetAllUsers(h.Database()),
	})
}

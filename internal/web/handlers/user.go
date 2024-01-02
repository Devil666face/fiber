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
	if h.View().IsHtmx() {
		return h.Render(view.UserContent, view.Map{
			view.UsersKey: models.GetAllUsers(h.Database()),
		})
	}
	return h.Render(view.UserList, view.Map{
		view.UsersKey: models.GetAllUsers(h.Database()),
	})
}

func UserEditForm(h *Handler) error {
	u := models.User{}
	id, err := strconv.Atoi(h.Ctx().Params("id"))
	if err != nil {
		return fiber.ErrNotFound
	}
	u.ID = uint(id)
	if err := u.Get(h.Database()); err != nil {
		return fiber.ErrNotFound
	}
	u.Password = ""
	return h.Render(view.UserEdit, view.Map{
		view.UserKey: u,
	})
}

func UserCreateForm(h *Handler) error {
	return h.Render(view.UserCreate, view.Map{})
}

func UserCreate(h *Handler) error {
	u := models.User{}
	if err := h.Ctx().BodyParser(&u); err != nil {
		return fiber.ErrBadRequest
	}
	if err := u.Validate(h.Validator()); err != nil {
		return h.Render(view.UserCreate, view.Map{
			view.UserKey:    u,
			view.MessageKey: err.Error(),
		})
	}
	if err := u.Create(h.Database()); err != nil {
		return h.Render(view.UserCreate, view.Map{
			view.UserKey:    u,
			view.MessageKey: err.Error(),
		})
	}
	return h.Render(view.UserCreate, view.Map{
		view.SuccessKey: fmt.Sprintf("User %s - created", u.Email),
	})
}

func UserEdit(h *Handler) error {
	var (
		u  = models.User{}
		in = models.User{}
	)
	if err := h.Ctx().BodyParser(&in); err != nil {
		return err
	}
	id, err := strconv.Atoi(h.Ctx().Params("id"))
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
			return h.Render(view.UserEdit, view.Map{
				view.UserKey:    u,
				view.MessageKey: err.Error(),
			})
		}
	}
	u.Email, u.Admin, u.Password = in.Email, in.Admin, in.Password
	if err := u.Update(h.Database()); err != nil {
		return h.Render(view.UserEdit, view.Map{
			view.UserKey:    u,
			view.MessageKey: err.Error(),
		})
	}
	if err := h.DestroySessionByID(u.SessionKey); err != nil {
		return h.Render(view.UserEdit, view.Map{
			view.UserKey:    u,
			view.MessageKey: err.Error(),
		})
	}
	return h.Render(view.UserEdit, view.Map{
		view.UserKey:    u,
		view.SuccessKey: "Successful update user",
	})
}

func UserDelete(h *Handler) error {
	u := models.User{}
	if err := h.Ctx().BodyParser(&u); err != nil {
		return err
	}
	id, err := strconv.Atoi(h.Ctx().Params("id"))
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
	return h.Render(view.UserContent, view.Map{
		view.UsersKey: models.GetAllUsers(h.Database()),
	})
}

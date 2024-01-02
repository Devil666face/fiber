package routes

import (
	"github.com/Devil666face/fiber/internal/web/handlers"
	"github.com/Devil666face/fiber/internal/web/middlewares"
)

func (r *Router) setUser() {
	user := r.router.Group("/user")
	user.Use(r.wrapper(middlewares.Auth))
	user.Use(r.wrapper(middlewares.Admin))

	user.Get(
		"/list",
		r.wrapper(handlers.UserListPage),
	).Name("user_list")

	user.Get(
		"/create",
		r.wrapper(middlewares.HxOnly),
		r.wrapper(handlers.UserCreateForm),
	).Name("user_create")
	user.Post(
		"/create",
		r.wrapper(middlewares.HxOnly),
		r.wrapper(handlers.UserCreate),
	)

	user.Get(
		"/:id<int>/edit",
		r.wrapper(middlewares.HxOnly),
		r.wrapper(handlers.UserEditForm),
	).Name("user_edit")
	user.Put(
		"/:id<int>/edit",
		r.wrapper(middlewares.HxOnly),
		r.wrapper(handlers.UserEdit),
	)

	user.Delete(
		"/:id<int>/delete",
		r.wrapper(middlewares.HxOnly),
		r.wrapper(handlers.UserDelete),
	).Name("user_delete")
}

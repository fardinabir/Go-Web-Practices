package users

import (
	"Go_CRUD_API/db/repos"
	"Go_CRUD_API/service"
	"github.com/go-chi/chi/v5"
)

type UserResource struct {
	Users UserStore
}

func NewResource() *UserResource {
	userStore := repos.NewUserStore()
	return &UserResource{Users: userStore}
}

func (rs *UserResource) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use()

	r.Mount("/users", rs.userRouter())
	r.Post("/login", rs.Login)
	r.Get("/", rs.HomePage)
	return r
}

func (rs *UserResource) userRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(service.NativeApiSecurity)
	r.Get("/{id:[0-9]+}", rs.ReadUser)
	r.Get("/", rs.ReadUsers)
	r.Post("/", rs.CreateUser)
	r.Patch("/{id:[0-9]+}", rs.UpdateUser)
	r.Delete("/{id:[0-9]+}", rs.DeleteUser)

	return r
}

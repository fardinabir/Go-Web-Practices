package server

import (
	"github.com/fardinabir/Go_CRUD_API/controllers/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New() (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	userResource := users.NewResource()
	r.Mount("/", userResource.Router())

	return r, nil
}

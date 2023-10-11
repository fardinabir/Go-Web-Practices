package server

import (
	"Go_CRUD_API/controllers/users"
	"Go_CRUD_API/database"
	"Go_CRUD_API/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
)

func New() (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	err := database.InitDatabase()
	if err != nil {
		log.Println("Postgres connection error", err)
	}

	err = service.InitRedisClient()
	if err != nil {
		log.Println("Redis connection error", err)
	}

	userResource := users.NewResource()
	r.Mount("/", userResource.Router())

	return r, nil
}

package users

import (
	"github.com/fardinabir/Go_CRUD_API/db/repos"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//type Controller struct {
//	DB *gorm.DB
//}

type UserResource struct {
	Users UserStore
}

func NewResource() *UserResource {
	userStore := repos.NewUserStore()
	return &UserResource{Users: userStore}
}

//func NewController() *Controller {
//	d := database.DatabaseConnection()
//	return &Controller{DB: d}
//}

func (rs *UserResource) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/", rs.userRouter())
	r.Post("/login", rs.Login)
	return r
}

func (rs *UserResource) userRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/users/{id:[0-9]+}", rs.ReadUser)
	r.Get("/users", rs.ReadUsers)
	r.Post("/users", rs.CreateUser)
	r.Delete("/users/{id:[0-9]+}", rs.DeleteUser)

	return r
}

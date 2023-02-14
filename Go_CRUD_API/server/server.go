package server

import (
	"Go_CRUD_API/controllers"
	"Go_CRUD_API/database"
	"log"

	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func init() {
	chi.RegisterMethod("PUT")
	chi.RegisterMethod("POST")
	chi.RegisterMethod("GET")
}

func StartServer(port string) {
	fmt.Println("Starting application at port : ", port)

	con := controllers.Controller{}
	con.DB = database.DatabaseConnection()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", controllers.HomePage)
	r.Get("/users/{id:[0-9]+}", controllers.ReadUser)
	r.Get("/users", controllers.ReadUsers)
	r.Post("/users", con.CreateUser)
	r.Delete("/users/{id:[0-9]+}", controllers.DeleteUser)

	r.Post("/login", con.Login)

	log.Fatal(http.ListenAndServe(":"+port, r))
}

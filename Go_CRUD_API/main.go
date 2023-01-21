package main

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

func main() {
	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", controllers.HomePage)
	r.Get("/users/{id:[0-9]+}", controllers.ReadUser)
	r.Get("/users", controllers.ReadUsers)
	r.Post("/users", controllers.CreateUser)
	r.Delete("/users/{id:[0-9]+}", controllers.DeleteUser)
	log.Fatal(http.ListenAndServe(":8085", r))

	//http.HandleFunc("/books", controllers.ReadBooks)
	//http.HandleFunc("")

	//r := gin.Default()
	//r.GET("/books/:id", controllers.ReadBook)
	//r.GET("/books", controllers.ReadBooks)
	//r.POST("/books", controllers.CreateBook)
	//r.PUT("/books/:id", controllers.UpdateBook)
	//r.DELETE("/books/:id", controllers.DeleteBook)
	//r.Run(":5000")
}

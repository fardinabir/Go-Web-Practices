package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter()

	// Initial Tasks
	myRouter.HandleFunc("/users", createNewUser).Methods("POST")
	myRouter.HandleFunc("/users/{id}", getUser).Methods("GET")

	// Extended Tasks
	myRouter.HandleFunc("/users/{id}/tags", createUserTag).Methods("POST")
	myRouter.HandleFunc("/users", getTagUsers).Methods("GET")

	// Others Utilities
	myRouter.HandleFunc("/allusers", returnAllUsers).Methods("GET")
	myRouter.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/users", updateUser).Methods("PUT")
	myRouter.HandleFunc("/alltags", returnAllTags).Methods("GET")

	http.Handle("/", myRouter)
	http.ListenAndServe(":8000", myRouter)
}

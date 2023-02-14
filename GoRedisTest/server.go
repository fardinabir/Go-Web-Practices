package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("On getUser..........")
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	fmt.Println(id)
	var tUser Student
	DB.First(&tUser, id)

	response, _ := json.Marshal(tUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}

func setUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("On setUser.......")
	var tUser Student
	json.NewDecoder(r.Body).Decode(&tUser)

	result := DB.Create(&tUser)
	if result.Error != nil {
		w.WriteHeader(500)
		fmt.Println(result.Error.Error())
		w.Write([]byte("Failed to insert..."))
		return
	}

	response, _ := json.Marshal("inserted")
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func Server() {
	r := chi.NewRouter()
	r.Get("/users/{id}", getUser)
	r.Post("/users", setUser)

	DatabaseConnection()
	log.Fatal(http.ListenAndServe(":8085", r))
}

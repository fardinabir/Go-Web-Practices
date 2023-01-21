package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"

	"Go_CRUD_API/database"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, this is the homepage")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("On CreateUser........")
	var tUser database.User
	json.NewDecoder(r.Body).Decode(&tUser)

	result := database.DB.Create(&tUser)
	fmt.Println(tUser) // temp lineeeeeeeeeeeeee
	if result.Error != nil {
		fmt.Println("Can't create the requested : ", result.Error)
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Server Error"})
		return
	}
	respondWithJSON(w, http.StatusCreated, map[string]string{
		"message":  "Created Successfully",
		"UserName": tUser.UserName,
		"id":       string(tUser.ID),
	})
}

func ReadUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("On ReadUser..........")
	var tUser database.User
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	result := database.DB.First(&tUser, id)
	if result.Error != nil {
		fmt.Println("Can't find the requested : ", result.Error)
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Server Error"})
		return
	}
	respondWithJSON(w, http.StatusOK, tUser)
}

func ReadUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("On ReadUsers........")
	var users []database.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		fmt.Println("Users not found")
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Server Error"})
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("On DeleteUsers.......")
	var tUser database.User
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	result := database.DB.Delete(&tUser, id)
	if result.Error != nil {
		fmt.Println("Delete failed, users not found")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Deleted Successfully"})
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

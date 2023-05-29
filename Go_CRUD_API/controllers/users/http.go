package users

import (
	"encoding/json"
	"fmt"
	"github.com/fardinabir/Go_CRUD_API/controllers"
	"github.com/fardinabir/Go_CRUD_API/model"
	"github.com/fardinabir/Go_CRUD_API/service"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"

	"github.com/fardinabir/Go_CRUD_API/database"
)

func (rs *UserResource) HomePage(w http.ResponseWriter, r *http.Request) {
	headerToken := service.GetHeaderValue(r, "Authorization")
	token, err := controllers.ValidateToken(headerToken)
	if err != nil {
		controllers.ErrUnauthorizedReq.ErrorResponse().JSONResponse(w)
		return
	}
	fmt.Println(token)
	fmt.Fprintf(w, "hello, this is the homepage")
}

func (rs *UserResource) ReadUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("On ReadUser..........")
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	res, err := rs.Users.GetUserById(id)
	if err != nil {
		log.Println("Can't find the requested : ", err.Error)
		controllers.ErrUserNotFound.ErrorResponse().JSONResponse(w)
		return
	}
	respondWithJSON(w, http.StatusOK, res)
}

func (rs *UserResource) ReadUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("On ReadUsers........")
	var users []model.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		fmt.Println("Users not found")
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Server Error"})
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

func (rs *UserResource) DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("On DeleteUsers.......")
	var tUser model.User
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

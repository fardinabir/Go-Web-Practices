package users

import (
	"Go_CRUD_API/controllers"
	"Go_CRUD_API/model"
	"Go_CRUD_API/service"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func (rs *UserResource) HomePage(w http.ResponseWriter, r *http.Request) {
	headerToken := service.GetHeaderValue(r, "Authorization")
	token, err := controllers.ValidateToken(headerToken)
	if err != nil {
		controllers.ErrUnauthorizedReq.ErrorResponse().JSONResponse(w)
		return
	}
	log.Println(token)
	fmt.Fprintf(w, "hello, this is the homepage")
}

func (rs *UserResource) ReadUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("On ReadUser..........")
	headerToken := service.GetHeaderValue(r, "Authorization")
	token, err := controllers.ValidateToken(headerToken)
	if err != nil {
		controllers.ErrUnauthorizedReq.ErrorResponse().JSONResponse(w)
		return
	}
	log.Println(token)

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
	headerToken := service.GetHeaderValue(r, "Authorization")
	token, err := controllers.ValidateToken(headerToken)
	if err != nil {
		controllers.ErrUnauthorizedReq.ErrorResponse().JSONResponse(w)
		return
	}
	log.Println(token)

	qry := map[string]interface{}{}
	users, err := rs.Users.GetUsers(qry) //database.DB.Find(&users)
	if err != nil {
		fmt.Println("Users not found")
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Server Error"})
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

func (rs *UserResource) UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("On UpdateUser.......")
	headerToken := service.GetHeaderValue(r, "Authorization")
	token, err := controllers.ValidateToken(headerToken)
	if err != nil {
		controllers.ErrUnauthorizedReq.ErrorResponse().JSONResponse(w)
		return
	}
	log.Println(token)

	var payload model.User
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	usr, err := rs.Users.UpdateById(id, &payload)
	if err != nil {
		log.Println("Update failed...", err)
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	log.Println("User Updated : ", usr)
	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Updated Successfully"})
}

func (rs *UserResource) DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("On DeleteUsers.......")
	headerToken := service.GetHeaderValue(r, "Authorization")
	token, err := controllers.ValidateToken(headerToken)
	if err != nil {
		controllers.ErrUnauthorizedReq.ErrorResponse().JSONResponse(w)
		return
	}
	log.Println(token)

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	usr, err := rs.Users.Delete(id)
	if err != nil {
		fmt.Println("Delete failed, users not found")
		controllers.ErrUserNotFound.ErrorResponse().JSONResponse(w)
		return
	}
	log.Println("User Deleted : ", usr)
	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Deleted Successfully"})
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

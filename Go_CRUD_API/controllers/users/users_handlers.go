package users

import (
	"Go_CRUD_API/controllers"
	"Go_CRUD_API/model"
	"Go_CRUD_API/service"
	"encoding/json"
	"fmt"
	"github.com/alexedwards/argon2id"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (rs *UserResource) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload model.User
	fmt.Println("This is CreateUser....")
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	_, err := rs.Users.GetUserByName(payload.UserName)
	if err == nil {
		controllers.ErrUserAlreadyExists.ErrorResponse().JSONResponse(w)
		return
	}

	hashPassword, err := argon2id.CreateHash(payload.Password, argon2id.DefaultParams)
	if err != nil {
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	payload.Password = hashPassword

	err = rs.Users.Create(&payload)
	if err != nil {
		fmt.Println("Can't create the requested : ", err.Error())
		controllers.ErrFailedToCreate.ErrorResponse().JSONResponse(w)
		return
	}

	// generating new token
	jwtAuth := controllers.NewTokenAuth()
	tokenResp := jwtAuth.GenerateTokens(payload.UserName)
	resp := &model.Response{Status: 200, Body: tokenResp}
	resp.JSONResponse(w)
}

func (rs *UserResource) Login(w http.ResponseWriter, r *http.Request) {
	var payload model.User
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	log.Print("This is payload", payload)

	// rate limiter using redis
	loginKey := "login" + payload.UserName
	rateLimitStatus, err := service.CheckRateLimit(loginKey, 3, time.Second*20)
	if rateLimitStatus == false {
		fmt.Println("---------------------Rate Limit Caught------------------------")
		controllers.ErrTooManyRequest.ErrorResponse().JSONResponse(w)
		return
	}

	userDb, err := rs.Users.GetUserByName(payload.UserName)
	if err != nil {
		controllers.ErrUserNotFound.ErrorResponse().JSONResponse(w)
		return
	}

	ok, err := argon2id.ComparePasswordAndHash(payload.Password, userDb.Password)

	if err != nil {
		controllers.ErrWrongPassword.ErrorResponse().JSONResponse(w)
		return
	}

	if ok { // generating new token
		jwtAuth := controllers.NewTokenAuth()
		tokenResp := jwtAuth.GenerateTokens(payload.UserName)
		resp := &model.Response{Status: 200, Body: tokenResp}
		resp.JSONResponse(w)
	}

}

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
	service.RespondWithJSON(w, http.StatusOK, res)
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
		service.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Server Error"})
		return
	}
	service.RespondWithJSON(w, http.StatusOK, users)
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
	service.RespondWithJSON(w, http.StatusOK, map[string]string{
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
	service.RespondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Deleted Successfully"})
}

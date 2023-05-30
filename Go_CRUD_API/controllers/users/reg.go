package users

import (
	"Go_CRUD_API/controllers"
	"Go_CRUD_API/model"
	"encoding/json"
	"fmt"
	"github.com/alexedwards/argon2id"
	"net/http"
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

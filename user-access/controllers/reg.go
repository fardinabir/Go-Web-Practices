package controllers

import (
	"Go_CRUD_API/model"
	"encoding/json"
	"fmt"
	"github.com/alexedwards/argon2id"
	"net/http"
)

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload, userDb model.User
	fmt.Println("This is CreateUser....")
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		fmt.Println(err)
		response := model.Response{Status: 500, Body: model.Error{Error: "Wrong Request Format"}}
		response.JSONResponse(w)
		return
	}

	err := c.DB.Where("user_name = ?", payload.UserName).First(&userDb).Error
	if err == nil {
		response := model.Response{Status: 500, Body: model.Error{Error: "User already exists!"}}
		response.JSONResponse(w)
		return
	}

	hashPassword, err := argon2id.CreateHash(payload.Password, argon2id.DefaultParams)

	if err != nil {
		response := model.Response{Status: 500, Body: err.Error()}
		response.JSONResponse(w)
		return
	}
	payload.Password = hashPassword

	err = c.DB.Create(&payload).Error
	if err != nil {
		fmt.Println("Can't create the requested : ", err.Error())
		response := model.Response{Status: 500, Body: err.Error()}
		response.JSONResponse(w)
		return
	}
	tokenResp := model.Token{AccessToken: "this is access token", RefreshToken: "this is refresh token"}
	response := model.Response{Status: 500, Body: tokenResp}
	response.JSONResponse(w)
}

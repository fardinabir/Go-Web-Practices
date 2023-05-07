package controllers

import (
	"Go_CRUD_API/model"
	"encoding/json"
	"fmt"
	"github.com/alexedwards/argon2id"
	"net/http"
)

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var payload, userDb model.User
	response := model.Response{Status: 500, Body: "Problem with Internal Server"}
	fmt.Println("This is Login....")
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		fmt.Println(err)
		response.JSONResponse(w)
		return
	}
	fmt.Println("This is payload", payload)
	err := c.DB.Where("user_name = ?", payload.UserName).First(&userDb).Error
	if err != nil {
		fmt.Println(err)
		response.JSONResponse(w)
		return
	}

	ok, err := argon2id.ComparePasswordAndHash(payload.Password, userDb.Password)

	if err != nil {
		response := model.Response{Status: 500, Body: "Wrong Password"}
		response.JSONResponse(w)
		return
	}

	if ok {

		tokenResp := model.Token{"this is access token", "this is refresh token"}
		response = model.Response{200, tokenResp}
		response.JSONResponse(w)
	}

}

func newFunc() {

}

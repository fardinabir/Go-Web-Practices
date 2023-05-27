package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/alexedwards/argon2id"
	"github.com/fardinabir/Go_CRUD_API/model"
	"log"
	"net/http"
)

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var payload, userDb model.User
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	log.Print("This is payload", payload)
	err := c.DB.Where("user_name = ?", payload.UserName).First(&userDb).Error
	if err != nil {
		fmt.Println(err)
		ErrUserNotFound.ErrorResponse().JSONResponse(w)
		return
	}

	ok, err := argon2id.ComparePasswordAndHash(payload.Password, userDb.Password)

	if err != nil {
		ErrWrongPassword.ErrorResponse().JSONResponse(w)
		return
	}

	if ok { // generating new token
		accToken := newToken(payload.UserName, 1, "access")
		refToken := newToken(payload.UserName, 15, "refresh")
		tokenResp := model.Token{AccessToken: accToken, RefreshToken: refToken}
		resp := &model.Response{Status: 200, Body: tokenResp}
		resp.JSONResponse(w)
	}

}

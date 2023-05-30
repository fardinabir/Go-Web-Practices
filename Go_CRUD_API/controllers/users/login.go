package users

import (
	"Go_CRUD_API/controllers"
	"Go_CRUD_API/model"
	"encoding/json"
	"github.com/alexedwards/argon2id"
	"log"
	"net/http"
)

func (rs *UserResource) Login(w http.ResponseWriter, r *http.Request) {
	var payload model.User
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		controllers.ErrInternalServerError.ErrorResponse().JSONResponse(w)
		return
	}
	log.Print("This is payload", payload)
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

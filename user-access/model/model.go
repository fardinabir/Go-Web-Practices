package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	UserName string `json:"userName"`
	MobileNo string `json:"mobileNo"`
	Password string `json:"password"`
}

type UserLogin struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type Error struct {
	Error string `json:"error"`
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Response struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

func (r *Response) JSONResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	json.NewEncoder(w).Encode(r.Body)
	return
}

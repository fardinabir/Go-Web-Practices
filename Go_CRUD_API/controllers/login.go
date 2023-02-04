package controllers

import (
	"Go_CRUD_API/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var payload, fromDb model.User
	fmt.Println("This is Login....")
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		response := model.Response{Status: 500, Body: model.Error{Error: "Wrong Request Format"}}
		response.JSONResponse(w)
	}
	fmt.Println(payload)
	fErr := c.DB.First(&fromDb, 1).Error
	if fErr != nil {
		fmt.Println(fromDb)
	}
}

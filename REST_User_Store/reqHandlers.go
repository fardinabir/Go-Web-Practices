package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func createNewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewUser")
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	fmt.Println(user)
	var res PostResponse
	res.ID = user.ID
	json.NewEncoder(w).Encode(res)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getUser")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	result := DB.First(&user, params["id"])
	fmt.Println(result)
	// Returning 404 Error
	if result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(404)
		return
	}
	var res GetResponse
	res.ID, res.Name, res.Phone = user.ID, user.FirstName+" "+user.LastName, user.Phone
	json.NewEncoder(w).Encode(res)
}

func createUserTag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateUserTag")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var req TagPostReq
	json.NewDecoder(r.Body).Decode(&req) // Decoding tags & expiry

	var user User
	result := DB.First(&user, params["id"])
	fmt.Println(result)
	// Catching Bad Requests
	if result.Error != nil {
		fmt.Println(result.Error)
		w.WriteHeader(400)
		return
	}

	var uTags []UserTag
	timeNowInMilli := (time.Now().UnixNano() / int64(time.Millisecond))
	for _, tag := range req.Tags {
		var uTag UserTag
		uTag.Tag = tag
		uTag.ExpiryTime = timeNowInMilli + req.Expiry // Expiry Time in Milliseconds
		uTags = append(uTags, uTag)
	}
	user.UserTags = uTags
	DB.Save(&user)
}

func getTagUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getTagUsers")
	w.Header().Set("Content-Type", "application/json")
	tags := regexp.MustCompile(` *, *`).Split(r.FormValue("tags"), -1)

	var uTags []UserTag
	DB.Where("tag IN ?", tags).Find(&uTags)
	var qIds []uint
	for _, uTag := range uTags {
		fmt.Println(uTag.UserID, uTag.ExpiryTime, (time.Now().UnixNano() / int64(time.Millisecond)))
		if uTag.ExpiryTime >= (time.Now().UnixNano() / int64(time.Millisecond)) {
			qIds = append(qIds, uTag.UserID)
			fmt.Println("ValidTime: ", uTag.UserID)
		}
	}

	var users []User
	DB.Model(&User{}).Preload("UserTags").Find(&users, qIds).Preload("UserTags") // Connected Ids

	var uTagResponses UserTagListResponse
	for _, user := range users {
		var uTagResponse UserTagResponse
		uTagResponse.ID, uTagResponse.Name = user.ID, user.FirstName+" "+user.LastName
		fmt.Println("Valid User : ", user.ID)
		for _, tagName := range user.UserTags {
			uTagResponse.Tags = append(uTagResponse.Tags, tagName.Tag)
		}
		uTagResponses.Users = append(uTagResponses.Users, uTagResponse)
	}
	fmt.Println(qIds)
	fmt.Println(len(qIds), len(uTagResponses.Users))
	json.NewEncoder(w).Encode(uTagResponses)
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllUsers")
	w.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users).Preload("UserTags")

	err := DB.Model(&User{}).Preload("UserTags").Find(&users).Error

	fmt.Println(err)
	json.NewEncoder(w).Encode(users)
}

func returnAllTags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllTags")
	w.Header().Set("Content-Type", "application/json")
	var uTags []UserTag
	DB.Find(&uTags)
	json.NewEncoder(w).Encode(uTags)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateUser")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	fmt.Println(user)
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteUser")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

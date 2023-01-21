package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	UserTags  []UserTag
}

type UserTag struct {
	gorm.Model
	UserID     uint
	Tag        string `json:"tag"`
	ExpiryTime int64  `json:"expiryTime"`
}

type GetResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"fullName"`
	Phone string `json:"phone"`
}

type PostResponse struct {
	ID uint `json:"id"`
}

type TagPostReq struct {
	Tags   []string `json:"tags"`
	Expiry int64    `json:"expiry"`
}

type UserTagListResponse struct {
	Users []UserTagResponse `json:users`
}

type UserTagResponse struct {
	ID   uint     `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

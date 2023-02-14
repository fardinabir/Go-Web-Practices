package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var err error

type Student struct {
	gorm.Model
	StdId   string `json:"stdId"`
	StdName string `json:"stdName"`
	//Subject interface{} `json:"subject"`
}

func dsnFetch() string {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "1234"
	dbname := "postgres"

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	return dsn
}

func DatabaseConnection() *gorm.DB {
	dsn := dsnFetch()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(Student{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
	return DB
}

//
//package database
//
//import (
//	"fmt"
//	"log"
//
//	"gorm.io/driver/postgres"
//	"gorm.io/gorm"
//)
//
//var DB *gorm.DB
//var err error
//
//type User struct {
//	gorm.Model
//	UserName string `json:"userName"`
//	MobileNo string `json:"mobileNo"`
//	Password string `json:"password"`
//}
//
//func DatabaseConnection() {
//	host := "localhost"
//	port := "5432"
//	dbName := "postgres"
//	dbUser := "postgres"
//	password := "1234"
//	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
//		host,
//		port,
//		dbUser,
//		dbName,
//		password,
//	)
//
//	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	DB.AutoMigrate(User{})
//	if err != nil {
//		log.Fatal("Error connecting to the database...", err)
//	}
//	fmt.Println("Database connection successful...")
//}

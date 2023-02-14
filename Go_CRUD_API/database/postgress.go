package database

import (
	"Go_CRUD_API/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var err error

func loadFromYaml() string {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file, ", err)
	}
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	return dsn
}

func DatabaseConnection() *gorm.DB {
	dsn := loadFromYaml()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(model.User{})
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

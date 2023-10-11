package database

import (
	"Go_CRUD_API/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func loadConfig() string {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	return dsn
}

func GetDBConnection() *gorm.DB {
	return db
}

func InitDatabase() error {
	if db != nil {
		return nil
	}
	var err error
	db, err = newDBConn()
	return err
}

func newDBConn() (*gorm.DB, error) {
	dsn := loadConfig()
	newDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = newDb.AutoMigrate(model.User{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Database connection successful...")
	return newDb, nil
}

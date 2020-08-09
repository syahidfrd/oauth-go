package common

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {

	dbHost := os.Getenv("DATABASE_HOST")
	dbName := os.Getenv("DATABASE_NAME")
	dbUserName := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")

	dbURI := fmt.Sprintf("host=%s dbname=%s user=%s  password=%s sslmode=disable", dbHost, dbName, dbUserName, dbPassword)

	db, err := gorm.Open("postgres", dbURI)

	if err != nil {
		fmt.Println("db error: ", err)
	}

	DB = db
	return DB

}

func GetDB() *gorm.DB {
	return DB
}

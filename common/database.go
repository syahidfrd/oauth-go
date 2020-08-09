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

	dbURI := os.Getenv("DATABASE_URL")
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

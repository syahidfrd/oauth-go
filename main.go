package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/subosito/gotenv"
	"gitlab.com/nyebarid/ny-rest-api/common"
	"gitlab.com/nyebarid/ny-rest-api/users"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
}

func main() {

	gotenv.Load()

	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	users.OAuthRegister(v1.Group("/auth"))

	v1.Use(users.IsAuth())
	users.SecretEndpointRegister(v1.Group("/secret"))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "up and running!",
		})
	})

	r.Run()
}

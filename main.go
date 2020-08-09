package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"gitlab.com/nyebarid/ny-rest-api/common"
)

func main() {

	gotenv.Load()

	db := common.Init()
	defer db.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "ny-rest-api up and running!",
		})
	})

	r.Run()
}

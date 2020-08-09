package users

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsAuth() gin.HandlerFunc {
	return checkJWT()
}

func checkJWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) != 2 {
			fmt.Println("passed")
			c.JSON(422, gin.H{
				"succes":  false,
				"message": "Provide a valid token",
			})
			c.Abort()
			return
		}

		token, _ := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("jwtUserID", claims["user_id"])
			c.Set("jwtIsAdmin", claims["user_role"])
		} else {
			c.JSON(422, gin.H{
				"succes":  false,
				"message": "Provide a valid token",
			})
			c.Abort()
			return
		}

	}
}

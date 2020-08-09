package users

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/danilopolani/gocialite.v1"
)

var gocial = gocialite.NewDispatcher()

func OAuthRegister(router *gin.RouterGroup) {
	router.GET("/:provider", OAuthRedirectHandler)
	router.GET("/:provider/callback", OAuthCallbackHandler)
}

func OAuthRedirectHandler(c *gin.Context) {
	provider := c.Param("provider")

	providerSecrets := map[string]map[string]string{
		"google": {
			"clientID":     os.Getenv("GOOGLE_CLIENT_ID"),
			"clientSecret": os.Getenv("GOOGLE_CLIENT_SECRET"),
			"redirectURL":  os.Getenv("GOOGLE_OAUTH_REDIRECT_URL") + "/google/callback",
		},
	}

	providerScopes := map[string][]string{
		"google": []string{},
	}

	providerData := providerSecrets[provider]
	actualScopes := providerScopes[provider]
	authURL, err := gocial.New().
		Driver(provider).
		Scopes(actualScopes).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	c.Redirect(http.StatusFound, authURL)
}

func OAuthCallbackHandler(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")
	provider := c.Param("provider")

	user, _, err := gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	userData, err := FindOneUser(&User{Provider: provider, SocialID: user.ID})

	if err != nil {
		newUser := User{
			Role:     "user",
			FullName: user.FullName,
			Email:    user.Email,
			Provider: provider,
			Avatar:   user.Avatar,
			SocialID: user.ID,
		}

		if err := SaveOne(&newUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Internal server error",
			})
			return
		}
		userData = newUser

	}

	jwtToken, err := createToken(&userData)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"user":         userData,
			"access_token": jwtToken,
		},
	})
}

func createToken(user *User) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"user_role": user.Role,
		"exp":       time.Now().AddDate(0, 0, 7).Unix(),
		"iat":       time.Now().Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("JWT_KEY")))

	return tokenString, err
}

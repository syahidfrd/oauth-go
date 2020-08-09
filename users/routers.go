package users

import (
	"fmt"
	"net/http"
	"os"

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

	user, token, err := gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	fmt.Printf("%#v", token)
	fmt.Printf("%#v", user)
	fmt.Printf("%#v", provider)

	c.Writer.Write([]byte("Hi, " + user.FullName))
}

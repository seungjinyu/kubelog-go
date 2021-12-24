package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {

	fmt.Println("Authenticating user")
	authToken := c.Request.Header.Get("auth-token")

	if authToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "No token",
		})
		return
	}
	if authToken != "secret-token" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		return
	}
	if len(c.Keys) == 0 {
		c.Keys = make(map[string]interface{})
	}
	c.Keys["received-token"] = authToken
	fmt.Println("Authenticating user completed")
	c.Next()

}

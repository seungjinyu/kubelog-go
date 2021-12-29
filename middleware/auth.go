package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/seungjinyu/kubelog-go/services"
)

type active struct {
	Active bool
}

func AuthenticationForBasic(c *gin.Context) {

	fmt.Println("Authenticating user")
	authToken := c.Request.Header.Get("auth-token")

	if authToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "No token",
		})
		c.Abort()
		return
	}
	if authToken != "secret-token" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		c.Abort()
		return
	}
	if len(c.Keys) == 0 {
		c.Keys = make(map[string]interface{})
	}
	c.Keys["received-token"] = authToken
	fmt.Println("Authenticating user completed")
	c.Next()

}

func AuthenticationForPod(c *gin.Context) {

	authorization := c.Request.Header.Get("authorization")
	authenticator := c.Request.Header.Get("authenticator")

	fmt.Println("authorizations", authorization)
	fmt.Println("authenticator", authenticator)

	if reflect.TypeOf(authenticator).String() == "string" && reflect.TypeOf(authorization).String() == "string" {
		fmt.Println("Split the bearer")

		token := strings.Split(authorization, "Bearer ")

		if token != nil {
			fmt.Println("Introspecting Token")

			data := services.IntrospectToken(token[1], authenticator)

			bdata, _ := json.Marshal(data)
			jsdata := active{}
			json.Unmarshal(bdata, &jsdata)
			if jsdata.Active {
				fmt.Println("Token authorized")
				c.JSON(http.StatusAccepted, gin.H{
					"result": jsdata.Active,
				})
				c.Next()

			} else {
				fmt.Println("Token not authorized")
				c.JSON(http.StatusUnauthorized, gin.H{
					"result": jsdata.Active,
				})
				c.Abort()
				return
			}

		}
		// code is still on development
	}

}

package auth

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/seungjinyu/kubelog-go/services"
)

func VerifyKey(c *gin.Context) {

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

			fmt.Println(data)
		}

	}

}

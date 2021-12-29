package services

import (
	"fmt"
	"os"

	"github.com/google/go-querystring/query"
	"github.com/monaco-io/request"
)

type customAuth struct {
	Username string `url:"client_id"`
	Password string `url:"client_secret"`
	Token    string `url:"token"`
}

func IntrospectToken(token string, authenticator string) interface{} {

	ca := customAuth{
		os.Getenv("AUTH_CLIENT_ID"),
		os.Getenv("AUTH_CLIENT_SECRET"),
		token,
	}
	v, _ := query.Values(ca)

	// fmt.Println("query: ", v)
	// fmt.Println("query encode: ", v.Encode())

	c := request.Client{
		URL:    "https://" + authenticator + os.Getenv("TOKEN_PATH"),
		Method: "POST",
		Header: map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
		// CustomerAuth: v.Encode(),
		String: v.Encode(),
		// Query: map[string]string{
		// 	"client_id":     os.Getenv("AUTH_CLIENT_ID"),
		// 	"client_secret": os.Getenv("AUTH_CLIENT_SECRET"),
		// 	"token":         token,
		// },
	}
	// fmt.Println(c.String)
	resp := c.Send()

	fmt.Println("resp", resp)

	return resp
}

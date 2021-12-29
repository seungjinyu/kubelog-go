package services

import (
	"fmt"
	"net/url"
	"os"
)

func IntrospectToken(token string, authenticator string) string {

	// qo = getQuery(token)

	// still on development

	values := url.Values{}
	values.Add("client_id", os.Getenv("AUTH_CLIENT_ID"))
	values.Add("client_secret", os.Getenv("AUTH_CLIENT_SECRET"))
	values.Add("token", token)
	query := values.Encode()

	fmt.Print(query)
	return ""
}

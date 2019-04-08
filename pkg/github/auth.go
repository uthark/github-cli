package github

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
)

// setAuthorization sets Authorization Header to request.
func (client Client) setAuthorization(req *http.Request) {
	authType := client.config.AuthType
	if authType == "token" {
		auth := base64.StdEncoding.EncodeToString([]byte(client.config.Token))
		req.Header.Add("Authorization", "Basic "+auth)
	} else if authType == "oauth2" {
		bytes, err := ioutil.ReadFile(client.config.AuthFile)
		if err != nil {
			log.Fatal("File with auth token doesn't exist or not accessible. login first: ", err)

		}
		req.Header.Add("Authorization", "token "+string(bytes))
	}
}

package login

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"

	"github.com/uthark/github-cli/pkg"
)

func GithubLogin(config *pkg.AppConfig) error {
	var conf *oauth2.Config
	var ctx context.Context

	localAddr := "localhost:33999"
	ctx = context.Background()
	conf = &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{"read:org"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("%s/login/oauth/authorize", config.GithubBaseURL),
			TokenURL: fmt.Sprintf("%s/login/oauth/access_token", config.GithubBaseURL),
		},
		RedirectURL: fmt.Sprintf("http://%s/oauth/callback", localAddr),
	}

	hc := &http.Client{}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, hc)

	authURL := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

	// Open Browser for auth.
	log.Println("Browser should be open for auth. If not, open link manually.")
	err := open.Run(authURL)
	if err != nil {
		log.Println("Unable to open browser", err)
	}

	log.Println("Authentication URL: ", authURL)

	// Start HTTP Server to wait for response from github.
	//http.HandleFunc("/oauth/callback", callbackHandler)
	h := http.NewServeMux()

	server := http.Server{Addr: localAddr}

	// Create channel to shutdown HTTP Server if Auth was successful.
	shutdown := make(chan bool, 1)
	go func() {
		<-shutdown
		defer func() {

			time.Sleep(1 * time.Second)
			err = server.Shutdown(ctx)
			if err != nil {
				log.Println(err)
			}
		}()
	}()

	callback := OauthCallback{
		shutdown: &shutdown,
		authFile: config.AuthFile,
		ctx:      ctx,
		conf:     conf,
	}
	h.HandleFunc("/oauth/callback", callback.ServeHTTP)

	server.Handler = h

	return server.ListenAndServe()
}

type OauthCallback struct {
	shutdown *chan bool
	authFile string
	conf     *oauth2.Config
	ctx      context.Context
}

func (c OauthCallback) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)
	if !strings.HasPrefix(r.RequestURI, "/oauth/callback") {
		return
	}

	queryParts, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Println(err)
	}

	// Use the authorization code that is pushed to the redirect URL.
	code := queryParts["code"][0]
	log.Println("Access Code:", code)

	token, err := c.conf.Exchange(c.ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Token:", token.AccessToken)

	err = ioutil.WriteFile(c.authFile, []byte(token.AccessToken), 0600)
	if err != nil {
		log.Fatal(err)
	}

	// The HTTP Client returned by conf.Client will refresh the token as necessary.
	client := c.conf.Client(c.ctx, token)

	resp, err := client.Get("https://api.github.com/events")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Authentication successful")

	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	// show success page
	msg := `<p><strong>Success!</strong></p>
	<p>You are authenticated and can now return to the CLI.</p>`

	_, err = fmt.Fprintf(w, msg)
	if err != nil {
		log.Println(err)
	}
	*c.shutdown <- true

}

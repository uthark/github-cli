package github

import (
	"io"
	"log"
	"net/http"
)

// closeWithErrorHandling will invoke close and log error if any.
func closeWithErrorHandling(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Println(err)
	}
}

// prepareRequest sets required Accept, User-Agent and Authorization headers.
func (client Client) prepareRequest(req *http.Request) {
	req.Header.Add("Accept", "application/vnd.github.inertia-preview.full+json")
	req.Header.Add("User-Agent", "uthark/github-cli")

	client.setAuthorization(req)
}

// min returns min of the numbers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

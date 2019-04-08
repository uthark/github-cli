package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/uthark/github-cli/pkg"
)

// Client is github API v3 HTTP Client. Does not implement full API.
type Client struct {
	config *pkg.AppConfig
}

// NewClient builds new client to interact with GitHub API.
func NewClient(config *pkg.AppConfig) *Client {
	return &Client{
		config: config,
	}
}

// ListOrganizationRepos lists all repos for the given organization.
func (client Client) ListOrganizationRepos(org string, projectCount int, sortType string) (RepositoriesWithPR, error) {
	log.Println("Using sort:", sortType)
	hc := http.Client{}
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/orgs/%s/repos", client.config.GithubAPIURL, org), nil)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build request")
	}

	client.prepareRequest(request)

	var repos RepositoriesWithPR
	// we will stop once we don't have extra results.
	log.Println("List of repos in organization:", org)

	for {
		resp, err := hc.Do(request)
		if err != nil {
			return nil, err
		}

		defer closeWithErrorHandling(resp.Body)

		if resp.StatusCode != http.StatusOK {
			decoder := json.NewDecoder(resp.Body)
			var apiError APIError
			err = decoder.Decode(&apiError)
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("organization not found or you don't have access to it: %s", apiError)
		}

		decoder := json.NewDecoder(resp.Body)

		link := resp.Header.Get("Link")
		var links = Links{}
		if link != "" {
			links, err = ParseLinks(link)
			if err != nil {
				return nil, err
			}

			// update query for the next page.
			query := request.URL.Query()
			query.Set("page", strconv.FormatInt(int64(links.NextPage), 10))

			request.URL.RawQuery = query.Encode()

		}
		var pageRepos []RepositoryWithPRs
		err = decoder.Decode(&pageRepos)
		if err != nil {
			return nil, err
		}

		repos = append(repos, pageRepos...)

		// if we don't have next page - stop.
		if links.NextPage == 0 {
			break
		}

	}

	return repos, nil

}

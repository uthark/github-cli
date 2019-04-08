package github

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// CountPullRequests counts pull requests for the given repo.
// owner - owner of the repo.
// repo - repo name.
// Since github API doesn't provide a simple method to determine number of pull requests,
// this method asks for PRs with per_page=1 and checks the number of returned pages in Link header.
func (client Client) CountPullRequests(owner, repo string) (int32, error) {

	hc := http.Client{}
	url := fmt.Sprintf("%s/repos/%s/%s/pulls?per_page=1&state=%s", client.config.GithubAPIURL,
		owner, repo, client.config.PRState)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, errors.Wrap(err, "cannot build request")
	}

	client.prepareRequest(request)

	resp, err := hc.Do(request)
	if err != nil {
		return 0, err
	}

	defer closeWithErrorHandling(resp.Body)

	if resp.StatusCode != http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		var apiError APIError
		err = decoder.Decode(&apiError)
		if err != nil {
			return 0, err
		}
		return 0, fmt.Errorf("repo not found or you don't have access to it: %s", apiError)
	}

	link := resp.Header.Get("Link")
	var links = Links{}
	if link != "" {
		links, err = ParseLinks(link)
		if err != nil {
			return 0, err
		}
	}

	return links.LastPage, nil
}

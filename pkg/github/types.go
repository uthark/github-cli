package github

import (
	"fmt"
	"strings"
	"time"
)

const (
	SortStars   = "stars"
	SortForks   = "forks"
	SortPRs     = "prs"
	SortContrib = "contrib"
)

// Repository is a github repository.
type Repository struct {
	ID        int32     `json:"id"`
	NodeID    string    `json:"node_id"`
	Name      string    `json:"name"`
	FullName  string    `json:"full_name"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Owner User `json:"owner"`

	ForksCount      int32 `json:"forks_count"`
	StargazersCount int32 `json:"stargazers_count"`
}

// RepositoryWithPRs is a repository with PR count attached.
// Decided to add a separate struct to distinguish "clean" response with added.
type RepositoryWithPRs struct {
	Repository
	PRCount int32
	contrib *float64
}

// Contrib calculates contributions. If there are no forks, contrib will return 0.
func (r RepositoryWithPRs) Contrib() float64 {
	if r.contrib == nil {
		var contrib = float64(0)
		if r.ForksCount != 0 {
			contrib = float64(r.PRCount) / float64(r.ForksCount)
		}
		r.contrib = &contrib
	}
	return *r.contrib
}

func (r RepositoryWithPRs) String() string {
	return fmt.Sprintf("%d\t%d\t%d\t%5.2f\t%s", r.ForksCount, r.StargazersCount, r.PRCount, r.Contrib(), r.FullName)
}

// User represents user in github.
type User struct {
	Login  string `json:"login"`
	ID     int32  `json:"id"`
	NodeID string `json:"node_id"`
	Type   string `json:"type"`
}

func (r Repository) String() string {
	return fmt.Sprintf("%d\t%d\t%s", r.ForksCount, r.StargazersCount, r.FullName)
}

// RepositoriesWithPR is a slice with repository with PRs.
type RepositoriesWithPR []RepositoryWithPRs

func (r RepositoriesWithPR) String() string {
	strs := make([]string, len(r))
	for i, v := range r {
		strs[i] = v.String()
	}
	return strings.Join(strs, "\n")
}

// APIError represents an error from GitHub API.
type APIError struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
}

func (p APIError) String() string {
	return p.Message
}

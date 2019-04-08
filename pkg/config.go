package pkg

import (
	"github.com/spf13/viper"
)

// AppConfig is an application config.
type AppConfig struct {
	// GithubAPIURL is Github API endpoint.
	GithubAPIURL string
	// GithubBaseURL is base URL to access Github.
	GithubBaseURL string
	// AuthType is an auth type used. Can be "token" or "oauth2"
	AuthType string
	// Token to access Github.
	Token string
	// OAuth2 Client ID
	ClientID string
	// OAuth2 Client Secret
	ClientSecret string
	// PRState defines which PRs include in filter.
	PRState string
	// AuthFile is a path to file with OAuth2 access token.
	AuthFile string
}

// NewAppConfig creates new AppConfig. Configuration is read from viper instance.
func NewAppConfig(viper *viper.Viper) *AppConfig {
	return &AppConfig{
		GithubAPIURL:  viper.GetString("global.api_url"),
		GithubBaseURL: viper.GetString("global.base_url"),
		AuthType:      viper.GetString("auth.type"),
		Token:         viper.GetString("auth.token"),
		ClientID:      viper.GetString("auth.client_id"),
		ClientSecret:  viper.GetString("auth.client_secret"),
		PRState:       viper.GetString("global.pr_state"),
		AuthFile:      viper.GetString("global.auth_file"),
	}

}

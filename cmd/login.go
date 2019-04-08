package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/uthark/github-cli/pkg"
	login2 "github.com/uthark/github-cli/pkg/github/login"
)

func login() *cobra.Command {
	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login to github",
		Long:  `Login to github and store the code for further usage.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			config := pkg.NewAppConfig(viper.GetViper())
			return login2.GithubLogin(config)
		},
		SilenceUsage: true,
		Example:      `  github-cli login`,
	}

	return loginCmd
}

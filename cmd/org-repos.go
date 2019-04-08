package cmd

import (
	"fmt"

	"github.com/uthark/github-cli/pkg"
	"github.com/uthark/github-cli/pkg/github"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// DefaultProjectCount is a number of projects to return by default.
	DefaultProjectCount = 20
	// DefaultSortOrder is default sort order for repos.
	DefaultSortOrder = "stars"
)

func orgRepos() *cobra.Command {
	count := DefaultProjectCount
	sort := DefaultSortOrder
	var orgReposCmd = &cobra.Command{
		Use:   "org-repos",
		Short: "List organization repositories",
		Long:  `List all repositories organization has.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("organization name not passed")
			}
			config := pkg.NewAppConfig(viper.GetViper())
			client := github.NewClient(config)

			return client.ShowStats(args[0], count, sort)
		},
		SilenceUsage: false,
		Example: `  github-cli org-repos netflix
  github-cli org-repos netflix -c 50 -s stars
  github-cli org-repos netflix -c 50 -s forks
  github-cli org-repos netflix -c 50 -s prs
  github-cli org-repos netflix -c 50 -s contrib
`,
	}

	orgReposCmd.Flags().IntVarP(&count, "count", "c", DefaultProjectCount,
		"Number of projects to return.")
	orgReposCmd.Flags().StringVarP(&sort, "sort", "s", DefaultSortOrder,
		"How to sort result. Valid values: stars, forks, prs, contrib")
	return orgReposCmd
}

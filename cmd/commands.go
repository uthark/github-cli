package cmd

import (
	"github.com/spf13/cobra"
)

type commandBuilder struct {
}

func (c commandBuilder) build() *cobra.Command {
	root := githubCli()

	c.addAll(root,
		orgRepos(),
		login(),
	)
	return root
}

func (c commandBuilder) addAll(parent *cobra.Command, subcommands ...*cobra.Command) {
	parent.AddCommand(subcommands...)
}

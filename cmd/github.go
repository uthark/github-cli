package cmd

import (
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// githubCli builds main command for CLI.
func githubCli() *cobra.Command {
	return &cobra.Command{
		Use:   "github-cli",
		Short: "Command-line tool to interact with Github.",
		Long: `Command-line tool to interact with Github.
                Source code available at https://github.com/uthark/github-cli`,

		PersistentPreRunE: loadConfig,
		SilenceErrors:     true,
	}
}

// loadConfig loads configuration using viper.
func loadConfig(_ *cobra.Command, _ []string) error {

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/github-cli/")
	viper.AddConfigPath("$HOME/.github-cli")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("GITHUB")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()

	log.Println("Loaded config: ", viper.ConfigFileUsed())
	if err != nil {
		return errors.Wrapf(err, "Fatal error config file: %s \n", viper.ConfigFileUsed())
	}
	return nil
}

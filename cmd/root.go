package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdRoot = &cobra.Command{
	Use:          "message",
	Short:        "Message sender with MessageBird API",
	SilenceUsage: true,
}

// Execute will perform the execute command given
func Execute() {
	cobra.OnInitialize(initConfig)
	if err := cmdRoot.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	err := viper.BindPFlags(cmdRoot.PersistentFlags())
	if err != nil {
		log.Fatal(err)
	}
}

// initConfig sets AutomaticEnv in viper to true.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
}

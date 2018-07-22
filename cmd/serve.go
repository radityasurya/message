package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	keyHost = "host"
	keyPort = "port"
)

var cmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Run message server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running")
	},
}

func init() {
	cmdRoot.AddCommand(cmdServe)
	cmdServe.Flags().StringP(keyHost, "H", "", "server host")
	cmdServe.Flags().IntP(keyPort, "p", 8080, "server port")
	err := viper.BindPFlags(cmdServe.Flags())
	if err != nil {
		log.Fatal(err)
	}
}

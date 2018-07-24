package cmd

import (
	"net/http"
	"time"

	"github.com/go-zoo/bone"
	"github.com/radityasurya/message/handler"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	keyHost           = "host"
	keyPort           = "port"
	keyMessageBirdAPI = "message-bird-key"
)

var cmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Run message server",
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString(keyMessageBirdAPI)
		if apiKey == "" {
			log.Fatalf("%s is not set", keyMessageBirdAPI)
		}

		h := viper.GetString(keyHost)
		p := viper.GetString(keyPort)

		r := bone.New()
		r.Get("/healthz", http.HandlerFunc(handler.Healthz))
		r.Post("/messages", http.HandlerFunc(handler.Message))

		srv := &http.Server{
			Handler:      r,
			Addr:         h + ":" + p,
			WriteTimeout: 30 * time.Second,
			ReadTimeout:  30 * time.Second,
		}

		log.Info("Running at ", h+":"+p)
		log.Fatal(srv.ListenAndServe())
	},
}

func init() {
	cmdRoot.AddCommand(cmdServe)

	cmdServe.Flags().StringP(keyHost, "H", "", "server host")
	cmdServe.Flags().IntP(keyPort, "p", 8080, "server port")
	cmdServe.Flags().StringP(keyMessageBirdAPI, "a", "", "messagebird access key")

	err := viper.BindPFlags(cmdServe.Flags())
	if err != nil {
		log.Fatal(err)
	}
}

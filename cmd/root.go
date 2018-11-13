package cmd

import (
	"fmt"
	"os"
	"time"

	"bitbucket.org/andyfusniakteam/capturoo-cli-go/configmgr"
	"github.com/spf13/cobra"
)

var (
	// Endpoint is a URL endpoint for the API
	Endpoint = "https://api.capturoo.com"
	Timeout  = time.Duration(10 * time.Second)
)

var caprc *configmgr.CapturooConfig

var rootCmd = &cobra.Command{
	Use:   "capturoo",
	Short: "capturoo is a tool for developers creating capturoo landing pages",
	Long:  `Complete documentation is available at https://www.capturoo.com`,
}

func init() {
	ep := os.Getenv("CAPTUROO_ENDPOINT")
	if ep != "" {
		Endpoint = ep
	}

	cobra.OnInitialize(initConfig)
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.capturoo.yaml)")
}

func initConfig() {
	caprc, _ = configmgr.ReadConfig()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

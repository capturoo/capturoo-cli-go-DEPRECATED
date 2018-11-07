package cmd

import (
	"fmt"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Endpoint is a URL endpoint for the API
	Endpoint = "https://api.capturoo.com"
	Timeout  = time.Duration(10 * time.Second)
)

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
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Search config in home directory with name ".capturoo" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".capturoo")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

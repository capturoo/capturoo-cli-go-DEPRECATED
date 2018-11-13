package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

// Account struct represents an Account resposne from the API
type Account struct {
	Aid           string `json:"aid"`
	Email         string `json:"email"`
	DisplayName   string `json:"displayName"`
	PrivateAPIKey string `json:"privateApiKey"`
	Created       string `json:"created"`
	LastModified  string `json:"lastModified"`
}

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Shows information about the capturoo account.",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Color("green")
		s.Start()

		tr := &http.Transport{
			MaxIdleConnsPerHost: 10,
		}

		client := &http.Client{
			Transport: tr,
			Timeout:   Timeout,
		}

		uri := fmt.Sprintf("%s/account", Endpoint)
		req, err := http.NewRequest("GET", uri, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\rerror creating new GET request: %v", err)
			s.Stop()
			os.Exit(1)
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Set("X-API-Key", caprc.PrivApiKey)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\rerror executing HTTP GET to %v : %v", uri, err)
			s.Stop()
			os.Exit(1)
		}
		defer resp.Body.Close()

		var account Account
		if err := json.NewDecoder(resp.Body).Decode(&account); err != nil {
			fmt.Fprintf(os.Stderr, "\rerror decoding url %s: %v\n", uri, err)
			s.Stop()
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "\rAccount ID: %s\n", account.Aid)
		fmt.Fprintf(os.Stdout, "Email: %s\n", account.Email)
		fmt.Fprintf(os.Stdout, "Created %s\n", account.Created)
		s.Stop()
	},
}

func init() {
	accountsCmd.AddCommand(infoCmd)
}

package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// Project struct represents a Project resposne from the API
type Project struct {
	PID          string `json:"pid"`
	ProjectName  string `json:"projectName"`
	LeadsCount   int    `json:"leadsCount"`
	PublicAPIKey string `json:"publicApiKey"`
	Created      string `json:"created"`
	LastModified string `json:"lastModified"`
}

// projectsListCmd represents the list command
var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows a list of capturoo projects.",
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

		uri := fmt.Sprintf("%s/projects", Endpoint)
		req, err := http.NewRequest("GET", uri, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\rerror creating new GET request: %v", err)
			s.Stop()
			os.Exit(1)
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Set("X-API-Key", "EbEweSE59l6u2SiLdgNdvYHj38oB1F1B0xYE149YTA2")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\rerror executing HTTP GET to %v : %v", uri, err)
			s.Stop()
			os.Exit(1)
		}
		defer resp.Body.Close()

		var plist []Project
		if err := json.NewDecoder(resp.Body).Decode(&plist); err != nil {
			fmt.Fprintf(os.Stderr, "\rerror decoding url %s: %v\n", uri, err)
			s.Stop()
			os.Exit(1)
		}

		s.Stop()

		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorder(false)
		table.SetHeader([]string{"Project ID", "Project Name", "Num Leads", "Public API Key"})
		for _, v := range plist {
			table.Append([]string{v.PID, v.ProjectName, strconv.Itoa(v.LeadsCount), v.PublicAPIKey})
		}
		table.Render()
	},
}

func init() {
	projectsCmd.AddCommand(projectsListCmd)
}

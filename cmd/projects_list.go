package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"bitbucket.org/andyfusniakteam/capturoo-cli-go/configmgr"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
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

// GetProjects makes a HTTP GET request to the capturoo API to get a slice of Projects.
func GetProjects(privApiKey string) ([]Project, error) {
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
		return nil, fmt.Errorf("error creating new GET request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-API-Key", privApiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error executing HTTP GET to %v : %v", uri, err)
	}
	defer resp.Body.Close()

	var plist []Project
	if err := json.NewDecoder(resp.Body).Decode(&plist); err != nil {
		return nil, fmt.Errorf("error decoding url %s: %v", uri, err)
	}

	return plist, nil
}

// projectsListCmd represents the list command
var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows a list of capturoo projects.",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Color("green")
		s.Start()

		pid, err := configmgr.ReadCurrentProject()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to read current project from $HOME/.capturoo/CURRENT_PROJECT: %v", err)
			os.Exit(1)
		}
		if pid == nil {
			fmt.Fprintf(os.Stderr, "failed to read current project from $HOME/.capturoo/CURRENT_PROJECT")
			os.Exit(1)
		}

		plist, err := GetProjects(caprc.PrivApiKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\rerror calling GetProjects: %v", err)
			s.Stop()
			os.Exit(1)
		}

		s.Stop()

		c := color.New(color.FgCyan)
		c.Add(color.Bold)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorder(false)
		table.SetHeader([]string{"Project ID", "Project Name", "Num Leads", "Public API Key"})
		for _, v := range plist {
			func(tpid string) {
				if tpid == *pid {
					table.Append([]string{c.Sprintf("%s", v.PID), v.ProjectName, strconv.Itoa(v.LeadsCount), v.PublicAPIKey})
				} else {
					table.Append([]string{v.PID, v.ProjectName, strconv.Itoa(v.LeadsCount), v.PublicAPIKey})
				}
			}(v.PID)
		}
		table.Render()
	},
}

func init() {
	projectsCmd.AddCommand(projectsListCmd)
}

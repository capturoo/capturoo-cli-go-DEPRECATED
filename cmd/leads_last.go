package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// leadsLastCmd represents the last command
var leadsLastCmd = &cobra.Command{
	Use:   "last",
	Short: "Retrieves the last lead added for the current project.",
	Run: func(cmd *cobra.Command, args []string) {
		if format != "json" && format != "csv" {
			fmt.Fprintf(os.Stderr, "Error: Format must be either json or csv.\n")
			os.Exit(1)
		}

		fmt.Println("last called")
		fmt.Println("Project is", projectName)
		fmt.Println("Format", format)
	},
}

func init() {
	leadsCmd.AddCommand(leadsLastCmd)
	leadsLastCmd.Flags().StringVarP(&projectName, "project-name", "p", "", "Project name")
}

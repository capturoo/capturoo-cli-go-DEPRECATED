package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var projectName string
var outfile string
var format string

// leadsListCmd represents the list command
var leadsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieves a list of leads for the current project.",
	Run: func(cmd *cobra.Command, args []string) {
		if format != "json" && format != "csv" {
			fmt.Fprintf(os.Stderr, "Error: Format must be either json or csv.\n")
			os.Exit(1)
		}

		fmt.Println("leads list called")
		fmt.Println("Project is", projectName)
		fmt.Println("Outfile", outfile)
		fmt.Println("Format", format)
	},
}

func init() {
	leadsCmd.AddCommand(leadsListCmd)

	leadsListCmd.Flags().StringVarP(&projectName, "project-name", "p", "", "Project name")
	leadsListCmd.Flags().StringVarP(&outfile, "output", "o", "", "Output filename")
	leadsListCmd.Flags().StringVarP(&format, "format", "f", "json", "Format json or csv")
}

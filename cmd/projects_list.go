package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// projectsListCmd represents the list command
var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows a list of capturoo projects.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("projecs list called")
	},
}

func init() {
	projectsCmd.AddCommand(projectsListCmd)
}

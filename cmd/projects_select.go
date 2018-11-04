package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Selects a project context.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("projects select called")
	},
}

func init() {
	projectsCmd.AddCommand(selectCmd)
}

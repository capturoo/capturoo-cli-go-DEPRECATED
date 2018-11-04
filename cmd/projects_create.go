package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new capturoo project.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	projectsCmd.AddCommand(createCmd)
}

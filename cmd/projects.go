package cmd

import (
	"github.com/spf13/cobra"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "The projects command can be used with in conjunction with the create, list and select subcommands.",
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}

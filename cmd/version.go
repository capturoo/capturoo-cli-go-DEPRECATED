package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print Capturoo version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, "%s\n", Version)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

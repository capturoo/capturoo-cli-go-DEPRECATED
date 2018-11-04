package cmd

import (
	"github.com/spf13/cobra"
)

// accountsCmd represents the accounts command
var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "The accounts command is used in conjuction with the info subcommand.",
}

func init() {
	rootCmd.AddCommand(accountsCmd)
}

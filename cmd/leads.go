package cmd

import (
	"github.com/spf13/cobra"
)

// leadsCmd represents the leads command
var leadsCmd = &cobra.Command{
	Use:   "leads",
	Short: "The leads command can be used in conjuction with the list or last subcommands.",
}

func init() {
	rootCmd.AddCommand(leadsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// leadsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// leadsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

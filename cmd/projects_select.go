package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"bitbucket.org/andyfusniakteam/capturoo-cli-go/configmgr"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Selects a project context.",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Color("green")
		s.Start()
		plist, err := GetProjects()
		if err != nil {
			fmt.Fprintf(os.Stderr, "\rerror calling GetProjects: %v", err)
			s.Stop()
			os.Exit(1)
		}
		s.Stop()

		// build a slice of "PID (Project)" strings
		pl := make([]string, 0, 50)
		for _, p := range plist {
			pn := fmt.Sprintf("%s (%s)", p.PID, p.ProjectName)
			pl = append(pl, pn)
		}

		sel := promptSelectProject(pl)
		words := strings.Fields(sel)
		pid := words[0]

		fmt.Fprintf(os.Stdout, "Project %s selected.\n", pid)
		err = configmgr.WriteCurrentProject(pid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to write currently selected project to filesystem: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	projectsCmd.AddCommand(selectCmd)
}

func promptSelectProject(pl []string) string {
	proj := ""
	prompt := &survey.Select{
		Message: "Choose a color:",
		Options: pl,
	}
	survey.AskOne(prompt, &proj, nil)
	return proj
}

package main

import (
	"bitbucket.org/andyfusniakteam/capturoo-cli-go/cmd"
)

var version string

func main() {
	cmd.Version = version
	cmd.Execute()
}

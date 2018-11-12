package configmgr

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

const (
	capturooFile = ".capturoorc"
	capturooDir  = ".capturoo"
)

func homeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return usr.HomeDir, nil
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func ensureCapturooDirExists() (string, error) {
	hd, err := homeDir()
	if err != nil {
		return "", err
	}

	capturooDir := filepath.Join(hd, capturooDir)

	exists, err := exists(capturooDir)
	if err != nil {
		return "", err
	}

	if !exists {
		os.Mkdir(capturooDir, 0755)
	}

	return capturooDir, nil
}

// WriteCurrentProject records the project ID on the filesystem within the $HOME/.capturoo directory in a file called CURRENT_PROJECT. The current project context can be read between invocation of the command-line tool.
func WriteCurrentProject(PID string) error {
	projectDir, err := ensureCapturooDirExists()
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't ensure capturoo dir exists: %s", err)
		os.Exit(1)
	}

	cpf := filepath.Join(projectDir, "CURRENT_PROJECT")
	bs := []byte(PID)
	err = ioutil.WriteFile(cpf, bs, 0644)
	if err != nil {
		return fmt.Errorf("failed to write CURRENT_PROJECT file: %v", err)
	}

	return nil
}

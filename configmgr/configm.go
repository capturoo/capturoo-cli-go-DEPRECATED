package configmgr

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/go-ini/ini"
)

const (
	capturooFile = ".capturoorc"
	capturooDir  = ".capturoo"
)

type CapturooConfig struct {
	AccountId  string
	PrivApiKey string
}

func homeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("failed user.Current(): %v", err)
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

func ensureCapturooDirExists() (*string, error) {
	hd, err := homeDir()
	if err != nil {
		return nil, fmt.Errorf("failed homeDir(): %v", err)
	}

	capturooDir := filepath.Join(hd, capturooDir)

	exists, err := exists(capturooDir)
	if err != nil {
		return nil, fmt.Errorf("failed exists(%s): %v", capturooDir, err)
	}

	if !exists {
		os.Mkdir(capturooDir, 0755)
	}

	return &capturooDir, nil
}

func ReadConfig() (*CapturooConfig, error) {
	hd, err := homeDir()
	if err != nil {
		return nil, err
	}

	cfg, err := ini.Load(filepath.Join(hd, capturooFile))
	if err != nil {
		return nil, fmt.Errorf("Fail to read file: %v", err)
	}

	dft := cfg.Section("default")
	caprc := CapturooConfig{
		AccountId:  dft.Key("account_id").String(),
		PrivApiKey: dft.Key("private_api_key").String(),
	}

	return &caprc, nil
}

// ReadCurrentProject read the $HOME/.capturoo/CURRENT_PROJECT file and returns the context. The file contains a single string that is the project ID.
func ReadCurrentProject() (*string, error) {
	projectDir, err := ensureCapturooDirExists()
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't ensure capturoo dir exists: %s", err)
		os.Exit(1)
	}

	cpf := filepath.Join(*projectDir, "CURRENT_PROJECT")
	bs, err := ioutil.ReadFile(cpf)
	if err != nil {
		return nil, err
	}

	pid := string(bs)
	return &pid, nil
}

// WriteCurrentProject records the project ID on the filesystem within the $HOME/.capturoo directory in a file called CURRENT_PROJECT. The current project context can be read between invocation of the command-line tool.
func WriteCurrentProject(PID string) error {
	projectDir, err := ensureCapturooDirExists()
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't ensure capturoo dir exists: %s", err)
		os.Exit(1)
	}

	cpf := filepath.Join(*projectDir, "CURRENT_PROJECT")
	bs := []byte(PID)
	err = ioutil.WriteFile(cpf, bs, 0644)
	if err != nil {
		return fmt.Errorf("failed to write CURRENT_PROJECT file: %v", err)
	}

	return nil
}

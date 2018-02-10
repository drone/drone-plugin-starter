package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

// Plugin defines the PyPi plugin parameters
type Plugin struct {
	Repository    string
	Username      string
	Password      string
	SetupFile     string
	Distributions []string
}

func (p Plugin) createConfig() error {
	f, err := os.Create(path.Join(os.Getenv("HOME"), ".pypirc"))
	if err != nil {
		return err
	}
	defer f.Close()
	buff := bufio.NewWriter(f)
	err = p.writeConfig(buff)
	if err != nil {
		return err
	}
	buff.Flush()
	return nil
}

func (p Plugin) writeConfig(buff io.Writer) error {
	_, err := io.WriteString(buff, fmt.Sprintf(`[distutils]
index-servers =
	repo

[repo]
repository: %s
username: %s
password: %s
`, p.Repository, p.Username, p.Password))
	return err
}

// buildCommand builds the exec.Command args to perform
// the desired python upload command
func (p Plugin) buildCommand() *exec.Cmd {
	// Set the default of distributions in here
	// as CLI package still has issues with string slice defaults
	distributions := []string{"sdist"}
	if len(p.Distributions) > 0 {
		distributions = p.Distributions
	}
	args := []string{p.SetupFile}
	for i := range distributions {
		args = append(args, distributions[i])
	}
	args = append(args, "upload")
	args = append(args, "-r")
	args = append(args, "repo")
	return exec.Command("python3", args...)
}

// Exec runs the plugin - doing the necessary setup.py modifications
func (p Plugin) Exec() error {
	err := p.createConfig()

	if err != nil {
		log.Fatalf("Unable to write .pypirc file due to: %s", err)
	}

	command := p.buildCommand()
	out, err := command.CombinedOutput()

	if err != nil {
		log.Fatalf("Error enountered: %s", out)
	}
	log.Printf("Output: %s", out)
	return nil
}

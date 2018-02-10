package main

import (
	"os"
	"strings"
	"testing"
)

func TestPublish(t *testing.T) {
	repository := os.Getenv("PLUGIN_REPOSITORY")
	username := os.Getenv("PLUGIN_USERNAME")
	password := os.Getenv("PLUGIN_PASSWORD")
	plugin := Plugin{
		Repository:    repository,
		Username:      username,
		Password:      password,
		SetupFile:     "testdata/setup.py",
		Distributions: strings.Split(os.Getenv("PLUGIN_DISTRIBUTIONS"), " "),
	}
	err := plugin.Exec()
	if err != nil {
		t.Error(err)
	}
}

// TestUpload checks if a distutils upload command can be properly
// generated and formatted.
func TestUpload(t *testing.T) {
	testdata := []struct {
		distributions []string
		exp           []string
	}{
		{
			[]string{},
			[]string{"python3", "testdata/setup.py", "sdist", "upload", "-r", "repo"},
		},
		{
			[]string{"sdist", "bdist_wheel"},
			[]string{"python3", "testdata/setup.py", "sdist", "bdist_wheel", "upload", "-r", "repo"},
		},
	}
	for i, data := range testdata {
		p := Plugin{Distributions: data.distributions, SetupFile: "testdata/setup.py"}
		c := p.buildCommand()
		if len(c.Args) != len(data.exp) {
			t.Errorf("Case %d: Expected %d, got %d", i, len(data.exp), len(c.Args))
		}
		for i := range c.Args {
			if c.Args[i] != data.exp[i] {
				t.Errorf("Case %d: Expected %s, got %s", i, strings.Join(data.exp, " "), strings.Join(c.Args, " "))
			}
		}
	}
}

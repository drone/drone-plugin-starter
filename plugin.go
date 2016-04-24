package main

type (
	Repo struct {
		Owner   string
		Name    string
		Link    string
		Avatar  string
		Branch  string
		Private bool
		Trusted bool
	}

	Build struct {
		Number   int
		Event    string
		Status   string
		Deploy   string
		Created  int64
		Started  int64
		Finished int64
		Link     string
	}

	Commit struct {
		Remote  string
		Sha     string
		Ref     string
		Link    string
		Branch  string
		Message string
		Author  Author
	}

	Author struct {
		Name   string
		Email  string
		Avatar string
	}

	Config struct {
		// plugin-specific parameters and secrets
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Commit Commit
		Config Config
	}
)

func (p Plugin) Exec() error {
	// plugin logic goes here
	return nil
}

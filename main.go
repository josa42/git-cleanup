package main

import (
	docopt "github.com/docopt/docopt-go"
	stringutils "github.com/josa42/go-stringutils"
)

func main() {
	usage := stringutils.TrimLeadingTabs(`
		Usage:
		  git-cleanup branches
		  git-cleanup keep     [--all]

		Options:
		  -h --help          Show this screen.
		  --version          Show version.
  `)

	arguments, _ := docopt.Parse(usage, nil, true, "Git Cleanup 0.1.0", false)

	cmdType := ""
	cmdTypes := []string{"home", "issues", "prs", "commits"}

	for _, key := range cmdTypes {
		if arguments[key] == true {
			cmdType = key
		}
	}

	switch cmdType {
	case "branches":
	case "keep":
	}
}

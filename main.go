package main

import (
	docopt "github.com/docopt/docopt-go"
	"github.com/josa42/git-cleanup/cleanup"
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

	arguments, _ := docopt.Parse(usage, nil, true, "Git Cleanup 0.1.1", false)

	cmdType := ""
	cmdTypes := []string{"branches", "keep"}

	for _, key := range cmdTypes {
		if arguments[key] == true {
			cmdType = key
		}
	}

	switch cmdType {
	case "branches":
		cleanup.Branches()
	case "keep":
		cleanup.Keep()
	}
}

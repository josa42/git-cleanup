package cleanup

import (
	"fmt"

	git "github.com/josa42/go-gitutils"
)

// Branches :
func Branches() {

	branch := git.CurrentBranch()
	mergedBranches := git.MergedBranches()

	fmt.Println("Clean branches")
	fmt.Println("=> Fetch")

	if branch == "master" && len(mergedBranches) > 0 {
		fmt.Println("=> Delete merged branches")
		for _, branch := range mergedBranches {
			fmt.Println("Delete:", branch)

			git.Exec("branch", "--delete", branch)
		}
	}
}

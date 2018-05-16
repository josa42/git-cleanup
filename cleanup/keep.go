package cleanup

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	git "github.com/josa42/go-gitutils"
	zglob "github.com/mattn/go-zglob"
)

// Keep :
func Keep() {
	fmt.Println("Clean .gitkeep files")
	keeps, _ := zglob.Glob("**/.gitkeep")

	for _, keep := range keeps {
		dir := path.Dir(keep)
		files, _ := filepath.Glob(dir + "/*")

		if len(files) > 1 && !git.IsIgnored(keep) {
			fmt.Println("=> Delete:", keep)
			os.Remove(keep)
		}
	}
}

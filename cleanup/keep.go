package cleanup

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	zglob "github.com/mattn/go-zglob"
)

// Keep :
func Keep() {
	keeps, _ := zglob.Glob("**/.gitkeep")

	for _, keep := range keeps {
		dir := path.Dir(keep)
		files, _ := filepath.Glob(dir + "/*")

		if len(files) > 1 {
			fmt.Println("Delete:", keep)
			os.Remove(keep)
		}
	}
}

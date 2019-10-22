package directory

import "path/filepath"

func Print(path string) string {
	return filepath.Base(path)
}

package plugins

import "path/filepath"

func PrintDirectory(path string) string {
	return filepath.Base(path)
}

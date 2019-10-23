package git

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"path/filepath"
)

const sigil = ""

func getRepo(path string) *git.Repository {
	if fi, err := os.Stat(filepath.Join(path, ".git")); os.IsNotExist(err) || !fi.IsDir() {
		return nil
	}

	r, err := git.PlainOpen(path)
	if err != nil {
		return nil
	}

	return r
}

func getBranch(r *git.Repository) string {
	ref, err := r.Head()
	if err != nil {
		return "HEAD"
	}

	return ref.Name().Short()
}

func Print(path string) string {
	r := getRepo(path)
	if r == nil {
		return ""
	}

	branch := getBranch(r)
	return fmt.Sprintf("%s %s", sigil, branch)
}

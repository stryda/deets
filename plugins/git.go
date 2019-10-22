package plugins

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
)

const sigill = ""

func getBranch(path string) string {
	r, err := git.PlainOpen(path)
	if err != nil {
		return ""
	}

	ref, err := r.Head()
	if err != nil {
		return "HEAD"
	}

	return ref.Name().Short()
}

func PrintGit(path string) string {
	branch := getBranch(path)
	if len(branch) > 0 {
		return fmt.Sprintf("%s %s", sigill, branch)
	}
	return ""
}

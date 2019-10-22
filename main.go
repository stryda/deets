package main

import (
	"fmt"
	"gitub.com/stryda/prompt/pkg/directory"
	"gitub.com/stryda/prompt/pkg/git"
	"log"
	"os"
	"strings"
)

func getDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func main() {
	currentDir := getDir()

	slugs := []string{
		directory.Print(currentDir),
		git.Print(currentDir),
	}

	output := strings.Join(slugs, " ")

	fmt.Println(output)
}

package main

import (
	plugin "./plugins"
	"fmt"
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
		plugin.PrintDirectory(currentDir),
		plugin.PrintGit(currentDir),
	}

	output := strings.Join(slugs, " ")

	fmt.Println(output)
}

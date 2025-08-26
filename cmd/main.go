package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	command := args[1]
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "No command provided")
		os.Exit(1)
	}

	switch command {
	case "init":
		CreateGitDirectory()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s", command)
		os.Exit(1)
	}
}

func CreateGitDirectory() {
	dir, _ := os.Getwd()
	path := filepath.Join(dir, ".git")
	os.Mkdir(path, 0700)
	path = filepath.Join(dir, ".git/objects")
	os.Mkdir(path, 0700)
	path = filepath.Join(dir, ".git/refs")
	os.Mkdir(path, 0700)
	data := []byte("ref: refs/heads/main\n")
	path = filepath.Join(dir, ".git/HEAD")
	os.WriteFile(path, data, 0700)
	fmt.Printf("Initialized empty Git repository in %s.git/\n", path)
}


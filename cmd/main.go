package main

import (
	"fmt"
	"os"
	"path/filepath"
	
	"github.com/btwkevin/kit/command"
)

func main() {
	args := os.Args
	cmd := args[1]
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "No cmd provided")
		os.Exit(1)
	}

	switch cmd {
	case "init":
		CreateGitDirectory()
	case "cat-file":
		command.HandleCatFileCommand()
	case "hash-object":
		HandleHashObjectCommand()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s", cmd)
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

// 9da8d6211e77c8a4dd43d929a769ded98592ec92
func HandleHashObjectCommand (){
	
}
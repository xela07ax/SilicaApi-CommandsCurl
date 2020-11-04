package main

import (
	billy "github.com/go-git/go-billy/v5"
	memfs "github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	memory "github.com/go-git/go-git/v5/storage/memory"
	"os"
)

var storer *memory.Storage
var fs billy.Filesystem

func main() {
	storer = memory.NewStorage()
	fs = memfs.New()
	_, err = git.PlainClone(directory, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
}
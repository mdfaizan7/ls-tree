package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var lastDir string

func createPath(path string, space int) {
	if path[0] == '.' || path == "node_modules" {
		return
	}
	paths := strings.SplitN(path, "\\", 2)
	if len(paths) <= 1 && lastDir != paths[0] {
		lastDir = paths[0]
		fmt.Println(strings.Repeat("│  ", space) + "├─ " + paths[0])
	} else {
		createPath(paths[1], space+1)
	}
}

func walkfunc(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		if f.Name() == ".git" || f.Name() == "node_modules" {
			return filepath.SkipDir
		}
	}
	createPath(path, 0)

	return nil
}

func GetTree() {
	filepath.Walk(".", walkfunc)
}

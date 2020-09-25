package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkfunc(path string, info os.FileInfo, err error) error {
	if path[0] != '.' {
		fmt.Println(path)
	}
	return nil
}

func main() {
	filepath.Walk(".", walkfunc)
}

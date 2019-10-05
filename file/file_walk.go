package main

import (
	"os"
	"path/filepath"
	_"log"
	"fmt"
)

func fileList(dir string) error {
	fullPath, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	
	callback := func(path string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(fullPath, path)
		if err != nil {
			return err
		}
		fmt.Println("rel:", rel, "\tbasepath:", fullPath, "\ttargpath:", path)
		return nil
	}

	return filepath.Walk(fullPath, callback)
}

func main(){
	fileList(".")
}

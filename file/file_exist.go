package main

import (
	"fmt"
	"os"
)

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println(err)
	return err == nil || os.IsExist(err)
}

func main(){
	if len(os.Args) > 1 {
		fmt.Println(FileExist(os.Args[1]))
	}
}



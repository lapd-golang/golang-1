package main
import (
	"fmt"
	"os"
	"log"
	"io"
	_"strings"
)

func main(){
	file, err := os.OpenFile("file.txt", os.O_APPEND | os.O_RDWR | os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	if _, err := file.Write([]byte("appended some data\n")); err != nil {
		log.Fatal(err)
	}
	if _, err := file.Seek(0, os.SEEK_SET); err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	err = nil
	for err == nil {
		n, err := file.Read(buf)
		if err == io.EOF {
			break	
		}
		fmt.Printf("Read %d bytes:", n)
		//fmt.Println(strings.TrimSpace(string(buf)))
		fmt.Println((string(buf)))
	}
}


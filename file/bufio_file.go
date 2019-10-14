package main
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main(){
	inputFile, inputError := os.Open("test.txt")
	if inputError != nil {
		fmt.Printf("An error occurred on openning th inputfile\n")
		return
	}

	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return 
		}
		fmt.Printf("The input was: %s", inputString)
	}
}

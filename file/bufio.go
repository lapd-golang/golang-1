package main
import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	var inputReader *bufio.Reader
	var input string
	var err error

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter some input:")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}

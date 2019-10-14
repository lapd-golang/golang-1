package main
import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Parameters:", os.Args[1:])
	fmt.Println("Parameters:", os.Args)
}

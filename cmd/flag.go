package main
import (
	"fmt"
	"flag"
)

func main(){
	var new_line = flag.Bool("n", false, "new line")
	var max_num int
	flag.IntVar(&max_num, "MAX_NUM", 100, "the num max")

	flag.PrintDefaults()
	flag.Parse()

	fmt.Println("There are", flag.NFlag(), "remaining args, they are:", flag.Args())
	fmt.Println("n has value: ", *new_line)
	fmt.Println("MAX_NJUM has value: ", max_num)
}

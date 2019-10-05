package main
import (
	"fmt"
)

func main() {
	set := map[int]string {}
	if set != nil {
		fmt.Println(set)
	} else {
		set = make(map[int]string)
		set[0]="hello"
		fmt.Println(set)

	}

	var set1  map[int] interface {}
	if set1 != nil {
		fmt.Println(set1)
	} else {
		set1 = make(map[int]interface{})
		set1[0]="hello"
		set1[2]=100
		fmt.Println(set1)
	}
}

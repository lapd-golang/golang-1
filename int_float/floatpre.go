package main

import (
	"fmt"
)

func main(){
	var a float32 =		1.0000001
	var b float32 =		1.00000001
	var bb float32 =	1.00000005
	var bbb float32 =	1.00000006
	var c float32 =		1.00000000000001

	fmt.Println(a == b) // false
	fmt.Println(a == bb) // false
	fmt.Println(b == bb) // true
	fmt.Println(b == c) // true	
	fmt.Println(bb == c)// true	
	fmt.Println(a == bbb) // true
	fmt.Println(bb == bbb) // false 
}

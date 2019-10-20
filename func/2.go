package main

import "fmt"

func multiParams(nums ...int) int {
	op := 0
	for _, i := range nums{
		op += i
	}
	return op
}

func main(){
	fmt.Println(multiParams(1,2,3))
	fmt.Println(multiParams(1,2,3, 4, 5))
	fmt.Println(multiParams(1))
	fmt.Println(multiParams())
	
}

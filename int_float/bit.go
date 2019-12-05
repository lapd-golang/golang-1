package main

import "fmt"

func main(){
	fmt.Printf("0x%X\n", 0xFF^0x55)
	var a uint8 = 0x55 
	fmt.Printf("%b\n", ^a)
}

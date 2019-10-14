package main

import "fmt"

func printCharsAndBytes(s string){
	for index, hao := range s {
		fmt.Printf("%c starts at byte %d\n", hao, index)
	}
}

func main(){
	name := "Señor"
	printCharsAndBytes(name)
}

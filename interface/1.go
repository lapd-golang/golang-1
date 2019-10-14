package main

import "fmt"

func assert(i interface{}){
	s,ok := i.(int)
	fmt.Println(s,ok)
}

func main(){
	var s interface{} = 56
	assert(s)
	
	var ss interface{} = "Hello" 
	assert(ss)
}

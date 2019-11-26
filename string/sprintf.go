package main

import (
	"fmt"
)

var base string = "%s is %d years old"
var second string = "&at %s\n"
func main(){
	s := fmt.Sprintf(base+"&at %s&", "wang", 10, "Beijing")
	fmt.Println(s)
}

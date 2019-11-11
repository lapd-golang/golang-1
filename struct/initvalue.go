package main

import (
	"fmt"
)

type People struct {
	Name string 
	Phone string
}

var p *People

func main(){
	if p == nil {
		fmt.Println("pointer initvalue is null")
	} else {
		fmt.Println(p)
	} 

	peo := new(People)
	if peo == nil {
		fmt.Println("struct initvalue is null")
	} else {
		fmt.Println(peo)
		peo.Name = "wang"
		peo.Phone = "123456"
		fmt.Println(peo)
	}
}

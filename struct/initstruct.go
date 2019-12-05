package main

import (
	"fmt"
)

type People struct {
	Name string 
	Phone string
	Addr []string
}

func main(){
	var pStruct People
//	if pStruct == nil {
//		fmt.Println("var initvalue is null")
//	} else {
		fmt.Println("var initvalue init: ",pStruct)
		fmt.Println("var initvalue init: Name: ", pStruct.Name)
		fmt.Println("var initvalue init: Name: ", pStruct.Phone)
		fmt.Println("var initvalue init: Name: ", pStruct.Addr, ", len:", len(pStruct.Addr))
		pStruct =  People{"wang", "123456", []string{"Beijing", "Shanghai"}}	
		fmt.Println("after var initvalue init: ", pStruct)
//	}
	fmt.Println("-------------------------------------------")

	var p *People
	if p == nil {
		fmt.Println("pointer initvalue is null")
	} else {
		fmt.Println("pointer initvalue is ", p)
	} 

	fmt.Println("-------------------------------------------")
	peo := new(People)
	if peo == nil {
		fmt.Println("new struct initvalue is null")
	} else {
		fmt.Println("new struct initvalue is ", peo)
		peo.Name = "wang"
		peo.Phone = "123456"
		fmt.Println("after new struct initvalue is ", peo)
	}
	fmt.Println("-------------------------------------------")

	ptmp, _ := testReturnPointer() 
	fmt.Printf("return pointer [%p]:[%p] value: %v\n", &ptmp, ptmp, *ptmp)
	fmt.Println("-------------------------------------------")

	tmp, _ := testReturn() 
	fmt.Printf("return [%p] value: %v\n", &tmp, tmp)
	fmt.Println("-------------------------------------------")
}
func testReturnPointer()(*People, error){
	tmp := People{} 
	tmp.Name = "wang"
	tmp.Phone = "123"
	fmt.Printf("In function, return pointer [%p] value: %v\n", &tmp, tmp)
	return &tmp, nil
}

func testReturn()(People, error){
	tmp := People{} 
	tmp.Name = "wang"
	tmp.Phone = "123"
	fmt.Printf("In function, return [%p] value: %v\n", &tmp, tmp)
	return tmp, nil
}

package main

import (
	"fmt"
	"strconv"
)

func main(){
	var ai [10]int
	var as [10]string
	fmt.Printf("int array [%p] init value: %v\n", &ai, ai)
	fmt.Printf("string array [%p] init value: %v\n", &as, as)
	fmt.Println("---------------------------")

	var i int
	for i = 0; i < len(ai); i++ {
		ai[i] = 100 + i
	}

	for i, _ := range as {
		as[i] = strconv.Itoa(100 + i)
	}
	fmt.Printf("int array [%p] init value: %v\n", &ai, ai)
	fmt.Printf("string array [%p] init value: %v\n", &as, as)
	fmt.Println("---------------------------")

	var pai *[10]int
	var pas *[10]string
	fmt.Printf("int point array [%p] init value: %v\n", &pai, pai)
	fmt.Printf("string point array [%p] init value: %v\n", &pas, pas)
	fmt.Println("---------------------------")

	pai = &ai
	pas = &as
	pai[0] = 1
	pas[0] = "1"
	fmt.Printf("int point array [%p] init value: %v\n", &pai, *pai)
	fmt.Printf("string point array [%p] init value: %v\n", &pas, *pas)
	fmt.Println("---------------------------")

	paii := new([10]int)
	pass := new([10]string)
	fmt.Printf("int point array [%p] init value: %v\n", &paii, *paii)
	fmt.Printf("string point array [%p] init value: %v\n", &pass, *pass)
	fmt.Println("---------------------------")

	for i = 0; i < len(*paii); i++ {
		paii[i] = 100 + i
	}
	for i, _ := range *pass {
		pass[i] = strconv.Itoa(100 + i)
	}
	fmt.Printf("int point array [%p] init value: %v\n", &paii, *paii)
	fmt.Printf("string point array [%p] init value: %v\n", &pass, *pass)
	fmt.Println("---------------------------")

	paii[0] = 1
	pass[0] = "1"
	fmt.Printf("%v\n==\n%v \n???\n%t\n", ai, *paii, (ai==*paii))
	fmt.Println("---------------------------")
	fmt.Printf("%v\n==\n%v \n???\n%t\n", as, *pass, (as==*pass))
}

package main

import (
	"fmt"
	"strconv"
)

func main(){
	var si []int
	var ss []string
	fmt.Printf("int slice [%p] init value: %v, len:%d, cap:%d\n", &si, si, len(si), cap(si))
	fmt.Printf("string slice [%p] init value: %v, len:%d, cap:%d\n", &ss, ss, len(ss), cap(ss))
//	fmt.Printf("int slice [%p] init value: %v\n", &si, si[0])
//	fmt.Printf("string slice [%p] init value: %v\n", &ss, ss[0])
	fmt.Println("---------------------------------------")

	var i int
	if si == nil {
		for i = 0; i < 10; i++ {
			si = append(si, 100+i)
			fmt.Printf("int slice [%p] append %d element: %v, len:%d, cap:%d\n", &si, i, si, len(si), cap(si))
		}
	}
	fmt.Println("---------------------------------------")
	if ss == nil {
		for i = 0; i < 10; i++ {
			ss = append(ss, strconv.Itoa(100+i))
			fmt.Printf("string slice [%p] append %d element: %v, len:%d, cap:%d\n", &ss, i, ss, len(ss), cap(ss))
		}
	}
	fmt.Println("---------------------------------------")
	
	var psi *[]int
	var pss *[]string
//	fmt.Printf("int slice pointer [%p] init value: %v, len:%d, cap:%d\n", &psi, psi, len(*psi), cap(*psi))
//	fmt.Printf("string slice pointer [%p] init value: %v, len:%d, cap:%d\n", &pss, pss, len(*pss), cap(*pss))
	fmt.Printf("int slice pointer [%p] init value: %v\n", &psi, psi)
	fmt.Printf("string slice pointer [%p] init value: %v\n", &pss, pss)
	fmt.Println("---------------------------------------")
	psi = &si
	pss = &ss
	fmt.Printf("int slice pointer [%p] init value: %v, len:%d, cap:%d\n", &psi, *psi, len(*psi), cap(*psi))
	fmt.Printf("string slice pointer [%p] init value: %v, len:%d, cap:%d\n", &pss, *pss, len(*pss), cap(*pss))
	fmt.Println("---------------------------------------")

	sim := make([]int, 10, 10)
	ssm := make([]string, 10, 10)
	fmt.Printf("int slice make [%p] init value: %v, len:%d, cap:%d\n", &sim, sim, len(sim), cap(sim))
	fmt.Printf("string slice make [%p] init value: %v, len:%d, cap:%d\n", &ssm, ssm, len(ssm), cap(ssm))
	fmt.Println("---------------------------------------")

	for i = 0; i < len(sim); i++ {
		sim[i] = 100 + i
	}
	for i,_ := range ssm {
		ssm[i] = strconv.Itoa(100+i)
	}
	fmt.Printf("int slice make [%p] init value: %v, len:%d, cap:%d\n", &sim, sim, len(sim), cap(sim))
	fmt.Printf("string slice make [%p] init value: %v, len:%d, cap:%d\n", &ssm, ssm, len(ssm), cap(ssm))
	fmt.Println("---------------------------------------")

//	fmt.Printf("%v\n==\n%v\n???\n%t\n", si, *psi, (si==*psi))
	fmt.Println("---------------------------------------")
//	fmt.Printf("%v\n==\n%v\n???\n%t\n", ss, ssm, (ss==ssm))

}

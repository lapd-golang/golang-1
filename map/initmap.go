package main

import (
	"fmt"
	"strconv"
)

func main(){
	var mi map[int]int
	var ms map[string]string
	fmt.Printf("int map [%p] init value: %v, len:%d\n", &mi, mi, len(mi))
	fmt.Printf("string map [%p] init value: %v, len:%d\n", &ms, ms, len(ms))
	fmt.Println("-------------------------------")

	var i int
	if mi == nil {
		mi = make(map[int]int)
		fmt.Printf("int map make [%p] init value: %v, len:%d\n", &mi, mi, len(mi))
		for i = 0; i < 10; i++ {
			mi[i] = 100 + i
			fmt.Printf("int map make [%p] init %d value: %v, len:%d\n", &mi, i, mi, len(mi))
		} 
	}
	fmt.Println("-------------------------------")

	if ms == nil {
		ms = make(map[string]string, 10)
		fmt.Printf("string map make [%p] init value: %v, len:%d\n", &ms, ms, len(ms))
		for i = 0; i < 10; i++ {
			k := strconv.Itoa(100+i)	
			ms[k] = strconv.Itoa(100+i)
			fmt.Printf("string map make [%p] init [%v] value: %v, len:%d\n", &ms, k, ms, len(ms))
		}
	}
	fmt.Println("-------------------------------")

	var pmi *map[int]int
	var pms *map[string]string
//	fmt.Printf("int map pointer [%p] init value: %v, len:%d\n", &pmi, pmi, len(*pmi))
//	fmt.Printf("string map pointer [%p] init value: %v, len:%d\n", &pms, pms, len(*pms))
	fmt.Printf("int map pointer [%p] init value: %v\n", &pmi, pmi)
	fmt.Printf("string map pointer [%p] init value: %v\n", &pms, pms)
	fmt.Println("-------------------------------")

	pmi = &mi
	pms = &ms
	fmt.Printf("int map pointer [%p] init value: %v, len:%d\n", &pmi, pmi, len(*pmi))
	fmt.Printf("string map pointer [%p] init value: %v, len:%d\n", &pms, pms, len(*pms))
	fmt.Println("-------------------------------")

//	fmt.Printf("%v\n==%v\n???\n%t\n", mi, *pmi, (mi==*pmi))
	fmt.Println("-------------------------------")
//	fmt.Printf("%v\n==%v\n???\n%t\n", ms, *pms, (ms==*pms))

	var mii map[int]int = map[int]int{}
	var mss map[string]string = map[string]string{}
	fmt.Printf("int map [%p] init value: %v, len:%d\n", &mii, mii, len(mii))
	fmt.Printf("string map [%p] init value: %v, len:%d\n", &mss, mss, len(mss))
	fmt.Println("-------------------------------")

	for i = 0; i < 10; i++ {
		mii[i] = 100 + i
	} 

	for i = 0; i < 10; i++ {
		k := strconv.Itoa(100+i)	
		mss[k] = strconv.Itoa(100+i)
	}
	fmt.Printf("int map [%p] init value: %v, len:%d\n", &mii, mii, len(mii))
	fmt.Printf("string map [%p] init value: %v, len:%d\n", &mss, mss, len(mss))
	fmt.Println("-------------------------------")
}

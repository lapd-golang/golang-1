package main

import (
	"time"
	"fmt"
)

func main(){
	var i, sum int

	st := time.Now()
	for i=0; i < 100000; i++{
		sum += i*i	
		time.Sleep(time.Microsecond)
	}
	et := time.Now()
	elapsed := et.Sub(st)
	fmt.Printf("Run time: %v\n", elapsed.Milliseconds())
	fmt.Println("The result is ", sum)
}

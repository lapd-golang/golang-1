package main

import (
	"fmt"
	"strconv"
	"errors"
)

func main(){
	slice, err := retSlice()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(slice)

	cases := [] struct{
		s []string
	}{
		s: {"I" ,"1" ,"2" ,"3" , "4" ,"5"},
//		{[]string{"I","1","2","3","4","5"}},
	//	{["I" "1" "2" "3" "4" "5"]},
	}

	fmt.Println(cases[0].s)
}

func retSlice()(args[]string, err error){
	args = make([]string, 6, 6)
	args[0]="I"
	args[1]="1"
	args[2]="2"
	args[3]="3"
	args[4]="4"
	args[5]=strconv.Itoa(19200)
	
	if len(args[5]) == 0 {
		err = errors.New("len is 0")
	} 

	fmt.Println(len(args))
	return 
}

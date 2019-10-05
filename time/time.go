package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	timeStr := now.Format("2000-10-01T10:30:30")
	timeStr1 := now.Format(time.ANSIC)
	fmt.Println("now:", now)
	fmt.Println("string:", timeStr)
	fmt.Println("string:", timeStr1)

	date , _ := time.Parse("2006-01-02 15:04:05", "2017-08-09 08:37:20")
	fmt.Println("string date:", date)

	trueOrFalse := date.After(now)
	if trueOrFalse == true {
		fmt.Println("After ", timeStr1)
	} else {
		fmt.Println("Before ", timeStr1)
	}
	
	m, _ := time.ParseDuration("-20m")
	m1 := now.Add(m)
	fmt.Println("Ten mintues:", m1)

}

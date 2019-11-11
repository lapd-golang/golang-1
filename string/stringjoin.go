package main

import (
	"fmt"
	"strconv"
)

func main(){
	errStr := "start:"
	var baudrate string
	baudrate = string(19200)

	errStr += "param" + strconv.Itoa(19200) + "\n"
	errStr += ", param" + baudrate + "\n"
	fmt.Printf(errStr)


}

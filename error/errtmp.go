package main

import (
	"fmt"
	"errors"
)

func errtmp(i int) (err error) {
//	err = fmt.Errorf(fmt.Sprintf("Protocol error"))
	if i > 5 {
		err = errors.New("Test error")
	}
	return err 
}

func main(){
	errt := errtmp(6)
	fmt.Println(errt.Error())
}


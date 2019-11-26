package main

import (
	"fmt"
	"errors"
	"unsafe"
)

func main(){
	var i int = 100
	fmt.Println(unsafe.Sizeof(i))
	fmt.Println(getTypeSize("REAL"))
}

func getTypeSize(str string) (int, error){
	switch str {
		case "bool", "BOOL":
			return 1, nil
		case "int8", "uint8", "byte", "SINT":
			return 1, nil
		case "int16", "uint16", "INT":
			return 2, nil
		case "int32", "uint32", "int", "uint", "DINT", "DWORD":
			return 4, nil
		case "int64", "uint64", "LINT":
			return 8, nil
		case "float32", "REAL":
			return 4, nil
		case "float64":
			return 8, nil
		default:
			return 0, errors.New("unknown type") 
	}
}

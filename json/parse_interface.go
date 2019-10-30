package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	b := []byte(`{"Name":"Wednesday", "Age":6, "Parents":["Gomez", "Morticia"]}`)
	
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		return
	}
	fmt.Println(":", f)

	if m, ok := f.(map[string]interface{}); ok {
		for k, v := range m {
			switch vv := v.(type){
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type that I don't know how to handle")
			}
		} 
	}
}

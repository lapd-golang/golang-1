package main

import (
	"fmt"
	sj "github.com/bitly/go-simplejson"
)

func main(){
	js, err := sj.NewJson([]byte(`{
		"test":{
			"array":[1, "2", 3],
			"int": 10,
			"float": 1.10,
			"bignum": 123456,
			"string": "simplejson",
			"bool": true
		}
	}`))
	if err != nil {
		return
	}

	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()

	fmt.Println("arr", arr, "i", i, "ms", ms)

}

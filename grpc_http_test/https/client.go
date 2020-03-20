package main

import (
	"fmt"
	_ "io/ioutil"
	"net/http"
	"time"
	"flag"
)

var count = flag.Int("count", 1000, "the connection times")

func main(){
	flag.Parse()

	var i int
	st := time.Now()

	for i=0; i < *count; i++{
		resp, err := http.Get("http://server:8088?numa=4&numb=6")
		if err != nil {
			fmt.Println("http get error: ", err)
			panic(err)
		}

		if resp.StatusCode != 200 {
			fmt.Println("http reponse error!")
		}

	//	body, _ := ioutil.ReadAll(resp.Body)
	//	fmt.Println(string(body))

		resp.Body.Close()
	}
	et := time.Now()
	elapsed := et.Sub(st)
	fmt.Printf("Run time: %v ms\n", elapsed.Milliseconds())
}

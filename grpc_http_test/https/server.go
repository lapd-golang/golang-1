package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var Addr string = ":8088"

func handler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
//	fmt.Println(r.Form)
//	fmt.Println(r.Form["numa"])
//	fmt.Println(r.Form["numb"])
	numa, _:= strconv.Atoi(r.Form["numa"][0])
	numb, _:= strconv.Atoi(r.Form["numb"][0])
	fmt.Fprintf(w, "%d\n", numa+numb)
}

func main(){
	http.HandleFunc("/", handler)
	
	fmt.Printf("listen...[%s]\n", Addr)
	err := http.ListenAndServe(Addr, nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

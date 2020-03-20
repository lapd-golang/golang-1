package main

import (
	"fmt"
	"net/http"
	"os"
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
	_, err := os.Open("cert/server.crt")
	if err != nil {
		fmt.Println("Can't open server.crt")
		panic(err)
	}

	fmt.Printf("listen...[%s]\n", Addr)
	err = http.ListenAndServeTLS(Addr, "cert/server.crt",
		"cert/server.key", nil)
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"os"
)

var Addr string = ":8088"

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, 
		" Hi, This is an example of https service in golang!\n")
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

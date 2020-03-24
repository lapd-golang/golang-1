package main

import (
	"fmt"
	"net/http"
	_"html"
	"log"
)

type fooHandler struct{
}

func (f fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, first!\n")
}

func zoo(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, zoo!\n")
}

func main(){
	http.Handle("/foo", fooHandler{})

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, %q\n", r.URL.Path)
	})

	http.Handle("/zoo", http.HandlerFunc(zoo))

	log.Fatal(http.ListenAndServe(":8088", nil))
}

package main

import (
	"log"
	"net/http"
)

func main(){
	srv := &http.Server{Addr: ":8088", Handler: http.HandlerFunc(handle)}

	log.Printf("Serving on http://server:8088")
	log.Fatal(srv.ListenAndServeTLS("cert/server.crt", "cert/server.key"))
}

func handle(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello"))
}



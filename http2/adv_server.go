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
	if r.URL.Path== "/2nd"{
		log.Println("Handling 2nd")
		w.Write([]byte("Hello Again!"))
		return
	}

	log.Println("Handling 1st")
	pusher, ok := w.(http.Pusher)
	if !ok {
		log.Println("can't push to client")
	} else {
		err := pusher.Push("/2nd", nil)
		if err != nil {
			log.Printf("Failed push: %v", err)
		}
	}
	w.Write([]byte("Hello"))
}



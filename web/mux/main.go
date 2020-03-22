package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func _Handler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions{
		return
	}
	w.Write([]byte("Gorilla!\n"))
}

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/foo", _Handler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	r.Use(mux.CORSMethodMiddleware(r))

	log.Fatal(http.ListenAndServe(":8000", r))
}

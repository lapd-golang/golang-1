package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"crypto/x509"
	"strconv"
)

type myhandler struct{
}

var Addr string = ":8088"

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
//	fmt.Println(r.Form)
//	fmt.Println(r.Form["numa"])
//	fmt.Println(r.Form["numb"])
	numa, _:= strconv.Atoi(r.Form["numa"][0])
	numb, _:= strconv.Atoi(r.Form["numb"][0])
	fmt.Fprintf(w, "%d\n", numa+numb)
}

func main(){
	pool := x509.NewCertPool()
	caCertPath := "cert/ca.crt"
	
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err: ", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr: Addr,
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs: pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	fmt.Println("listen...")
	err = s.ListenAndServeTLS("cert/server.crt", "cert/server.key")
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"crypto/x509"
)

type myhandler struct{
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, 
		" Hi, This is an example of https service in golang!\n")
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
		Addr: ":8088",
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

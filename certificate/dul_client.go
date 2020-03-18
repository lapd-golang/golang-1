package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"crypto/tls"
	"crypto/x509"
)

func main(){
	// x509.Certificate
	pool := x509.NewCertPool()

	caCertPath := "cert/ca.crt"
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("cert/client.crt", "cert/client.key")
	if err != nil {
		fmt.Println("LoadX509keypair err: ", err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}

	//resp, err := client.Get("https://localhost:8088")
	resp, err := client.Get("https://server:8088")
	if err != nil {
		fmt.Println("http get error: ", err)
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.Status)
}

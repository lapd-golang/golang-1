package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"time"
)

var count = flag.Int("count", 1000, "the connection times")

func main(){
	flag.Parse()

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

	var i int
	st := time.Now()

	for i=0; i < *count; i++{

		resp, err := client.Get("https://server:8088?numa=4&numb=6")
		//resp, err := client.Get("https://localhost:8088")
		if err != nil {
			fmt.Println("http get error: ", err)
			panic(err)
		}

		if resp.StatusCode != 200 {
			fmt.Println("http reponse error!")
		}
		//body, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(body))
//		fmt.Println(resp.Proto)

		resp.Body.Close()
		client.CloseIdleConnections()
	}
	et := time.Now()
	elapsed := et.Sub(st)
	fmt.Printf("Run time: %v ms\n", elapsed.Milliseconds())
}

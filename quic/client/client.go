package main

import (
	_ "bytes"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	_ "io"
	"net/http"
	"io/ioutil"
	"time"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

var count = flag.Int("count", 1000, "the connection times")
var addr = flag.String("addr", "https://server:8088?numa=4&numb=6", "connect to addr")
var insecure = flag.Bool("insecure", false, "skip certificate verification")

func main() {
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

	var qconf quic.Config

	var i int
	st := time.Now()
	for i=0; i < *count; i++{

		roundTripper := &http3.RoundTripper{
			TLSClientConfig: &tls.Config{
				RootCAs:            pool,
				InsecureSkipVerify: *insecure,
			},
			QuicConfig: &qconf,
		}
		hclient := &http.Client{
			Transport: roundTripper,
		}

		rsp, err := hclient.Get(*addr)
		if err != nil {
			fmt.Println("http3 get error: ", err)
			panic(err)
		}
		if rsp.StatusCode != 200 {
			fmt.Println("http reponse error!")
		}
/*
		fmt.Printf("Got response for %s: %#v", *addr, rsp)
		body := &bytes.Buffer{}
		_, err = io.Copy(body, rsp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Request Body: %s", body.Bytes())
*/		rsp.Body.Close()
		roundTripper.Close()
	}

	et := time.Now()
	elapsed := et.Sub(st)
	fmt.Printf("Run time: %v ms\n", elapsed.Milliseconds())
}

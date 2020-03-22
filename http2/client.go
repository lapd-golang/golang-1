package main

import (
	"fmt"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"golang.org/x/net/http2"
)

var addr = flag.String("addr", "https://server:8088?numa=4&numb=6", "connect to")
var httpVer = flag.Int("httpVer", 2, "HTTP version")

func main(){
	flag.Parse()

	client := &http.Client{}

	caCert, err := ioutil.ReadFile("cert/ca.crt")
	if err != nil {
		log.Fatalf("Reading server certificate: %s", err)
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)
	tlsConfig := &tls.Config{
		RootCAs: pool,
	}

	switch *httpVer {
		case 1:
			client.Transport = &http.Transport {
				TLSClientConfig: tlsConfig,
			}
		case 2:
			client.Transport = &http2.Transport {
				TLSClientConfig: tlsConfig,
			}
	}

	resp, err := client.Get(*addr)
	if err != nil {
		log.Fatalf("Failed get: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}

	fmt.Printf("Response %d: %s\nbody: %s\n", resp.StatusCode, resp.Proto, string(body))
}

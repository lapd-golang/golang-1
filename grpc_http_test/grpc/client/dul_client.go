/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"time"
	"path/filepath"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"fmt"
	"flag"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "example/helloworld"
)

const (
	//	address     = "https://server:8088"
	address     = "server:8088"
	defaultName = "world"
	certsDir = "cert/"
)

func getTransportCredentials() (*credentials.TransportCredentials, error) {
	crtPath := filepath.Clean(filepath.Join(certsDir, "client.crt"))
	keyPath := filepath.Clean(filepath.Join(certsDir, "client.key"))
	caPath := filepath.Clean(filepath.Join(certsDir, "ca.crt"))

	cert, err := tls.LoadX509KeyPair(crtPath, keyPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, err
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, errors.Errorf("Failed append CA certs from %s", caPath)
	}

	creds := credentials.NewTLS(&tls.Config{
//		ServerName: "server",
//		ClientAuth:	tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	})

	fmt.Println("crtPath: "+crtPath + ", key: "+ keyPath + ", ca: " + caPath)
	return &creds, nil
}
/*
func createConnection(ctx context.Context, address string) *grpc.ClientConn {
	tc, err := getTransportCredentials()
	if err != nil {
		fmt.Println("Error when creating transport credentials: " + err.Error())
		os.Exit(1)
	}

	conn, err := grpc.DialContext(ctx, address,
		grpc.WithTransportCredentials(*tc), grpc.WithBlock())

	if err != nil {
		fmt.Println("Error when dialing: " + address + " err:" + err.Error())
		os.Exit(1)
	}

	return conn
}
*/

var addr = flag.String("addr", "server:8088", "the address to connect to")
var count = flag.Int("count", 1000, "the connection times")

func main() {
	flag.Parse()

	// Set up a connection to the server.
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	fmt.Println("conn...")

	tc, err := getTransportCredentials()
	if err != nil {
		fmt.Println("Error when creating transport credentials: " + err.Error())
		os.Exit(1)
	}

	var i int
	st := time.Now()
	for i=0; i< *count; i++{
		conn, err := grpc.DialContext(ctx, *addr, grpc.WithTransportCredentials(*tc), grpc.WithBlock())
		if err != nil {
			fmt.Println("Error when dialing: " + address + " err:" + err.Error())
			os.Exit(1)
		}
		c := pb.NewGreeterClient(conn)

//		fmt.Println("call...")
		// Contact the server and print out its response.
//		r, err := c.SayHello(ctx, &pb.HelloRequest{Numa: "4", Numb: "6"})
		_, err = c.SayHello(ctx, &pb.HelloRequest{Numa: "4", Numb: "6"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
//		log.Printf("Greeting: %d\n", r.GetNumc())
		conn.Close()
	}
	et := time.Now()
	elapsed := et.Sub(st)

	log.Printf("time: %d ms\n", elapsed.Milliseconds())
}

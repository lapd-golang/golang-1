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
	"time"
	"flag"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "example/helloworld"
)

var addr = flag.String("addr", "server:8088", "the address to connect to")
var count = flag.Int("count", 1000, "the connection times")

func main() {
	flag.Parse()

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	creds, err := credentials.NewClientTLSFromFile("cert/ca.crt", "")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	var i int
	st := time.Now()
	for i=0; i< *count; i++{

		// Set up a connection to the server.
		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c := pb.NewGreeterClient(conn)

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

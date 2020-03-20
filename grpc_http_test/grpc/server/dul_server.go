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

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"flag"
	"strconv"
	
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "example/helloworld"
)

const (
	certsDir = "cert/"
)

var port = flag.Int("port", 8088, "the port to serve on")

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
//	log.Printf("Received: %v, %v", in.GetNuma(), in.GetNumb())
	numa, _ := strconv.Atoi(in.GetNuma())
	numb, _ := strconv.Atoi(in.GetNumb())
	return &pb.HelloReply{Numc: int32(numa+numb)}, nil
}

func getTransportCredentials() (*credentials.TransportCredentials, error) {
	crtPath := filepath.Clean(filepath.Join(certsDir, "server.crt"))
	keyPath := filepath.Clean(filepath.Join(certsDir, "server.key"))
	caPath := filepath.Clean(filepath.Join(certsDir, "ca.crt"))

	cert, err := tls.LoadX509KeyPair(crtPath, keyPath)
	if err != nil {
		return nil, fmt.Errorf("Failed load server key pair: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, fmt.Errorf("Failed appends CA certs from %s: %s", caPath, err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, errors.Errorf("Failed append CA certs from %s", caPath)
	}

	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:      certPool,
	})

	fmt.Println("crtPath: "+crtPath + ", key: "+ keyPath + ", ca: " + caPath)

	return &creds, nil
}

func main() {
	flag.Parse()

	tc, err := getTransportCredentials()
	if err != nil {
		log.Fatalf("fialed to read certificates: %s", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer(grpc.Creds(*tc))
	pb.RegisterGreeterServer(s, &server{})
	fmt.Println("listen ", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

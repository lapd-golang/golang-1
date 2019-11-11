package main

import (
	"log"
	"time"

	"github.com/goburrow/serial"
)

func main() {
	port, err := serial.Open(&serial.Config{Address: "com10"})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	_, err = port.Write([]byte("serial"))
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(10*time.Second)	
}

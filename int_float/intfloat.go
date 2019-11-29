package main

import (
	"io"
	"bytes"
	"encoding/binary"
	"fmt"
)

func encodeValue(value interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, value)
	return buf.Bytes(), err 
}

func decodeValue(reader io.Reader, value interface{}) error {
	err := binary.Read(reader, binary.LittleEndian, value)
	return err
}

func float32Value(buf []byte) (value float32, err error) {
	err = decodeValue(bytes.NewReader(buf), &value)
	return value, err
}

func int16Value(buf []byte) (value int16, err error) {
	err = decodeValue(bytes.NewReader(buf), &value)
	return value, err
}

func uint16Value(buf []byte) (value uint16, err error) {
	err = decodeValue(bytes.NewReader(buf), &value)
	return value, err
}

func main(){
	var ff float32 = 100.5
	ffb, err := encodeValue(ff)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ffb)
	}

	fmt.Println("------------------")
	ffn, err := float32Value(ffb)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ffn)
	}

	fmt.Println("----------------------------------")
	var ss int16 = -100 
	ssb , err := encodeValue(ss)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ssb)
	}
	fmt.Println("------------------")
	ssn, err := int16Value(ssb)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ssn)
	}

	fmt.Println("----------------------------------")
	var us uint16 = 0xFFF1 
	usb , err := encodeValue(us)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(usb)
	}
	fmt.Println("------------------")
	usn, err := uint16Value(usb)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("0x%X\n", usn)
		var mask uint16 = 0x1
		var maskv uint16 = 0x0
		usn = usn & (^mask) | mask 
		fmt.Printf("0x%X\n", usn)
		usn = usn & (^mask) | maskv 
		fmt.Printf("0x%X\n", usn)
	}
}

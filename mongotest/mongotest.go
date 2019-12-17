package main

import (
	"fmt"
	"github.com/yuxi-o/golang/mongotest/model"
	"time"
)

var vb bool = true

func main() {
	for {
		if vb {
			vb = false
		} else {
			vb = true
		}
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("----------------%v---------------%v------\n", time.Now().Unix(), vb)
		go query()
	}
}

func query() {
	conf, err := model.FindByNameDf1("DF1SS", vb)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("config: %+v\n", conf)

	sli, err := model.GetServerArgs(conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server: ", sli)

	slic, err := model.GetClientArgs(conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Client: ", slic)
}

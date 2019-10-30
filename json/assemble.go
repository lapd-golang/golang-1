package main

import (
	"encoding/json"
	"fmt"
)

type Server struct{
	ServerName string `json:"serverName"`
	ServerIP string `json:"serverIP"`
	Name string `json:"-"`
}

type Serverslice struct {
	Servers [] Server `json:"servers"`
	StrFirst string `json:"strFirst"`
	StrSecond string `json:"strSecond,string"`
}

func main(){
	var s Serverslice	
	s.StrFirst = `Go "1.0" ` 
	s.StrSecond = `Go "1.0" ` 
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1", Name:"wang"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}

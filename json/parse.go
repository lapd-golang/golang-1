package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName	string
	serverIP	string
	Name	string
	Addr	string
	Addr1	string
}

type Serverslice struct {
	Servers []Server
}

func main(){
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN", "serverIP":"127.0.0.1", "name":"wang"}, {"serverName":"Beijing_VPN", "serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}


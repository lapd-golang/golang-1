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

	fmt.Println("---------------Unmarshal-----------")
	str := `{"Name":"Wednesday", "Age":6, "Parents":["Gomez", "Morticia"]}`
	var f interface{}
	err = json.Unmarshal([]byte(str), &f)
	if err != nil {
		return
	}
	fmt.Println(":", f)

	if m, ok := f.(map[string]interface{}); ok {
		for k, v := range m {
			switch vv := v.(type){
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type that I don't know how to handle")
			}
		} 
	}

}

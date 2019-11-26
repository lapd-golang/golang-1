package main

import "fmt"

func change(s ...string){
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s, len(s))
}

func main(){
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome, len(welcome))

	here := []string{"I", "Love", "China"}
	fmt.Printf("%p:%v, len:%d, cap:%d\n", &here, here, len(here), cap(here))
	here[2]="Beijing"
	fmt.Printf("%p:%v, len:%d, cap:%d\n", &here, here, len(here), cap(here))
	mm := map[string]string{"addr":"A", "name":"B"}
	if str, ok := mm["addr"]; !ok {
		home := str + "false"
		fmt.Println(home)
	} else {
		home := str + "true"
		fmt.Println(home, mm)
	}

	remote := []string{"I", "Love", "China"}
	fmt.Printf("%p\n", remote)
	remotet := []string{"A", "B"}
	remote = append(remote[:1], remotet...)
	fmt.Println(remote, len(remote))
	fmt.Printf("%p\n", remote)
	ps := &remote
	fmt.Println(ps[0], ps[1])

	buf := []byte{'A', 'B', 'C', 'D'}
//	buf[-1] = 'E'
//	buf[-2] = 'F'
	fmt.Println(buf)

	buff := make([]byte, 3)
	fmt.Println(len(buff))
}

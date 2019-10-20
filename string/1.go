package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
	)

func main(){
	var s string
	fmt.Println(s)
	s = "hello"
	fmt.Println(len(s))
//	s[1] = 'w'
	s = "\xE4\xBA\xB5\xFF"
	fmt.Println(s)
	fmt.Println(len(s))
	s = "中"
	fmt.Println(s)
	fmt.Println(len(s))
	c := []rune(s)
	fmt.Printf("中 unicode %x\n", c[0])
	fmt.Printf("中 utf8 %x\n", s);

	ss := "中华人民共和国"
	for _, c := range ss {
		fmt.Printf("%c %x\n", c, c)
	}

	str := "A,B,C"
	parts := strings.Split(str, ",")
	for _, part := range parts {
		fmt.Println(part)
	}

	fmt.Println(strings.Join(parts, "-"))

	sd := strconv.Itoa(10)
	fmt.Println("str"+sd)
	if i, err := strconv.Atoi("10"); err == nil {
		fmt.Println(10+i)
	}

	os.Exit(0)
}

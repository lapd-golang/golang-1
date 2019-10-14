package main
import "fmt"

type Test interface {
	Tester()
}

type myFloat float64

func (m myFloat) Tester(){
	fmt.Println(m)
}

func describe(t Test){
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func main(){
	var t Test
	f := myFloat(89.7)
	t = f
	describe(t)
	t.Tester()
}

package main
import (
	"fmt"
	"os"
	"math/rand"
	"time"
)

func returnMultiValues() (int, int){
	t := time.Now()
	rand.Seed(t.Unix());
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int{
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int{
	time.Sleep(time.Second*1)
	return op
}

func main(){
	fmt.Println(returnMultiValues())	
	tsSF := timeSpent(slowFunc)
	fmt.Println(tsSF(10))

	os.Exit(0)
}

package main
import (
	"fmt"
	"time"
)

func finished(v int){
	fmt.Printf("Finished something:%d\n", v)
}

func process(v int){
	defer finished(v)
	fmt.Printf("Start processing...%d\n", v)
	time.Sleep(2*time.Second)	
	v = 3
	fmt.Printf("End processing...%d\n", v)
}

func main(){
	process(5)
}

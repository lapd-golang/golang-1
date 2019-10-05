package main
import (
	"fmt"	
	"os"
	"os/exec"
	"strings"
)

func main(){
	sliceFunc(os.Args[1:]...)
}

func sliceFunc(cmd... string){
	fmt.Println(cmd)
	if len(cmd) == 0 {
		fmt.Printf("Usage: %s args...\n", os.Args[0])
		os.Exit(-1)
	}
	fmt.Println(cmdFunc(cmd...))
}

func cmdFunc(cmd... string) string {
	fmt.Printf("cmd slice len: %d, value:%v\n", len(cmd),  cmd)
	result, err := exec.Command(cmd[0], cmd[1:]...).Output()
	if err != nil {
		fmt.Println("Command failed:", err.Error())
	}

//	return string(result)  // with '\n'
	return strings.TrimSpace(string(result))
}

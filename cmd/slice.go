package main
import (
	"fmt"	
	"os"
	"os/exec"
	"strings"
	"strconv"
	"time"
)

func main(){
	/*
	err := checkCmdStatus()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("check OK")
	}	
	*/
	//sliceFunc(os.Args[1:]...)

	cmdstr := make([]string, 2, 2)
	cmdstr[0] = os.Args[1]
	var i int = 0
	for {
		i++
		cmdstr[1] = strconv.Itoa(i)
		fmt.Println("starting 1...", cmdstr)
		go cmdFuncRun(cmdstr...)
		fmt.Println("starting 2...", cmdstr)
		time.Sleep(1*time.Second)
	}
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
func cmdFuncRun(cmd... string) error {
	fmt.Printf("cmd slice len: %d, value:%v\n", len(cmd),  cmd)
	err := exec.Command(cmd[0], cmd[1:]...).Run()
	if err != nil {
		fmt.Println("Command failed:", err.Error())
		return err
	}
	
	return nil
}
/*
func checkCmdStatus()(error){
	checkCmd := []string{"pgrep", "df1d"}
	
	err := exec.Command(checkCmd[0], checkCmd[1]).Run()
	return err
	 
}
*/

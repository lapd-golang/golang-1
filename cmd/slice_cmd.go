package main
import (
	"fmt"	
	"os"
	"os/exec"
	"strings"
	"strconv"
)

func main(){
	rwCmd := []string{"df1c", "127.0.0.1", "N7:1"}
	stopCmd := []string{"pkill", "df1d"}
	restartCmd := []string{"df1d"}

	// read or write data 
	str, err := cmdFunc(rwCmd...)
	if err != nil {
		fmt.Println(err)	// error code
		fmt.Println(string(str)) // error reason
	} else {
		fmt.Println(string(str))
	}

	// restart service after read or write failed more times
	if i, _ := strconv.Atoi(os.Args[1]); i == 3 {
		 exec.Command(stopCmd[0], stopCmd[1]).Run()
		 exec.Command(restartCmd[0], restartCmd[1:]...).Run()
	}
}

func cmdFunc(cmd... string) (string, error) {
	fmt.Printf("cmd len: %d, value:%v\n", len(cmd),  cmd)
	result, err := exec.Command(cmd[0], cmd[1:]...).Output()
	if err != nil {
		return strings.TrimSpace(string(result)), err 
	}

	return strings.TrimSpace(string(result)), nil
}



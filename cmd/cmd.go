package main

import (
	"fmt"
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main(){
	if len(os.Args) == 1{
		fmt.Printf("Usage: %s args...\n", os.Args[0])
		os.Exit(-1)
	}
	str1, err := RunCommandWithErr(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(str1)	
	}

	str := RunCommand(os.Args[1])	
	fmt.Println(str)	
}

func runInLinux(cmd string) string{
	fmt.Println("Running Linux cmd:" , cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result))
}

func runInWindows(cmd string) string{
	fmt.Println("Running Win cmd:", cmd)
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result))
}

func RunCommand(cmd string) string{
	if runtime.GOOS == "windows" {
		return runInWindows(cmd)
	} else {
		return runInLinux(cmd)
	}
}

func RunLinuxCommand(cmd string) string{
	if runtime.GOOS == "windows" {
		return ""
	} else {
		return runInLinux(cmd)
	}
}

func runInLinuxWithErr(cmd string) (string, error) {
	fmt.Println("Running Linux cmd:"+cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result)), err
}

func runInWindowsWithErr(cmd string) (string, error){
	fmt.Println("Running Windows cmd:"+cmd)
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result)), err
}

func RunCommandWithErr(cmd string) (string, error){
	if runtime.GOOS == "windows" {
		return runInWindowsWithErr(cmd)
	} else {
		return runInLinuxWithErr(cmd)
	}
}

func RunLinuxCommandWithErr(cmd string)(string, error){
	if runtime.GOOS == "windows" {
		return "", errors.New("could not run in Windows Os") 
	} else {
		return runInLinuxWithErr(cmd)
	}
}

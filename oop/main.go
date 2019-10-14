package main
import "github.com/yuxi-o/golang/oop/employee"

func main(){
	/*
	e := employee.Employee{
		FirstName: "wang",
		LastName: "qing",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	*/
	e := employee.New("wang", "qing", 30, 20)
	e.LeavesRemaining()
}

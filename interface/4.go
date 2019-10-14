package main

import "fmt"

type SalaryCalculator interface{
	DisplaySalary()
}

type LeaveCalculator interface{
	CalculateLeavesLeft() int
}

type EmployeeOperations interface{
	SalaryCalculator
	LeaveCalculator
}

type Employee struct {
	firstName string
	lastName string
	basicPay int
	pf int
	totalLeaves int
	leavesTaken int
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d\n", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main(){
	e := Employee{
		firstName: "Wang",
		lastName: "qing",
		basicPay: 5000,
		pf:200,
		totalLeaves:30,
		leavesTaken:5,
	}

	var empOp EmployeeOperations
	empOp = e
	empOp.DisplaySalary()
	fmt.Println("Leaves left = ", empOp.CalculateLeavesLeft())
}


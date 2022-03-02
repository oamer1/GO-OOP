package main

import "oop/employee"


func main(){
	e1 := employee.Employee {
		Name: "Adel Fahmy",
		Salary: 10000,
	}
	// What happens if Employee is not initialized ?
	// Strings default to "" and int defaults to 0
	var e2 = employee.Employee{}
	e1.GetSalary()
	e2.GetSalary()
}
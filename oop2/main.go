package main

import (
	"oop/employee"
)

func main() {
	e := employee.New("Adel Fahmy",10000)
	e.GetSalary()
}
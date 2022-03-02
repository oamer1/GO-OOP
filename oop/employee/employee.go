package employee

import "fmt"

// struct and its fields is available outside package
type Employee struct {
	Name   string
	Salary int
}

func (e Employee) GetSalary() {
	fmt.Printf("Employee \"%s\" Salary is %d\n",e.Name, e.Salary)
}

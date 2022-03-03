package employee

import "fmt"

// By using lowercase 'employee' instead of "Employee", we prevented access from other packages
// Expose it only through New method
// Also use lowercase for fields so they are unexported too
type employee struct {
	name   string
	salary int
}

// Factory function act as constructor
func New(name string, salary int) employee{
	e := employee {name, salary}
	return e
}

func (e employee) GetSalary() {
	fmt.Printf("Employee \"%s\" Salary is %d\n",e.name, e.salary)
}

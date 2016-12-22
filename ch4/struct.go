package main

import "time"
import "fmt"

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var dilbert Employee
	dilbert.Salary -= 5000

	position := &dilbert.Position
	*position = "Senior " + *position

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += "(proactive team player)"
	//Same as:
	(*employeeOfTheMonth).Position += "(proactive team player)"

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position)

	id := dilbert.ID
	EmployeeByID(id).Salary = 0
}

func EmployeeByID(id int) *Employee {
	/*...*/
	return &Employee{}
}

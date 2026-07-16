package employee

import "fmt"

type employee struct {
	id     int
	name   string
	salary float64
}

func (emp *employee) IncreaseSalary(percent float64) {
	fmt.Println("Current Salary:", emp.salary)
	fmt.Println("Increment Percentage:", percent*100, "%")
	emp.salary += (emp.salary * percent)
	fmt.Println("New Salary:", emp.salary)

}
func (emp *employee) PrintDetails() {
	fmt.Println("Employee Id:", emp.id)
	fmt.Println("Employee Name:", emp.name)
	fmt.Println("Employee Salary:", emp.salary)

}

func RegisterEmployee() {
	emp1 := employee{
		id:     1,
		name:   "Nishak",
		salary: 250000,
	}

	emp1.IncreaseSalary(0.50)
	emp1.PrintDetails()
}

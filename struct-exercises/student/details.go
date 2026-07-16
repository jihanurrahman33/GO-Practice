package student

import "fmt"

func printStudentDetails(std Student) {

	fmt.Println("Student Name:", std.name)
	fmt.Println("Student Age:", std.age)
	fmt.Println("Student DEPT:", std.department)
	fmt.Println("Student CGPA:", std.cgpa)
	fmt.Println()
}

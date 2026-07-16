package main

import "student_class_management/student"

func main() {

	std1 := student.Student{
		Name: "Jihan",
	}

	std1.AddMarks(81, 82, 83, 84, 85, 86, 87, 88, 89, 90)

	std1.Display()
}

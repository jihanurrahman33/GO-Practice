package main

import (
	"fmt"
	"struct_practice/bank"
	"struct_practice/car"
	"struct_practice/employee"
	"struct_practice/rectangle"
	"struct_practice/student"
	"struct_practice/students"
)

func main() {
	student.CallProblem1()
	rectangle.CallProblem2()
	fmt.Println()
	bank.RegisterBank()
	employee.RegisterEmployee()
	car.RegisterCar()
	students.RegisterStudents()
}

package students

import "fmt"

type student struct {
	name string
	roll int
	cgpa float64
}

var studentList = []student{}

func (std *student) addStudent() {
	studentList = append(studentList, *std)
	fmt.Println(studentList)
}

func displayStudentList() {
	// for _, std := range studentList {
	// 	fmt.Println(std)
	// }
	fmt.Println(studentList)
}
func findHighestCGPAStudent() {
	maxCGPAStudent := studentList[0]
	for idx, std := range studentList {
		if std.cgpa > maxCGPAStudent.cgpa {
			maxCGPAStudent = studentList[idx]

		}
	}
	fmt.Println("Max CGPA Student:", maxCGPAStudent)
}

func avgCGPA() {
	totalCGPA := 0.0
	for _, std := range studentList {
		totalCGPA += float64(std.cgpa)
	}
	avgCGPA := totalCGPA / float64(len(studentList))
	fmt.Printf("Average CGPA:%.2f", avgCGPA)
}

func RegisterStudents() {
	std1 := student{
		name: "Jihan",
		roll: 4969,
		cgpa: 3.19,
	}
	std2 := student{
		name: "Nishak",
		roll: 4970,
		cgpa: 3.20,
	}
	std3 := student{
		name: "Jihanur",
		roll: 4971,
		cgpa: 3.21,
	}
	std4 := student{
		name: "Jr",
		roll: 4972,
		cgpa: 3.22,
	}

	std1.addStudent()
	std2.addStudent()
	std3.addStudent()
	std4.addStudent()
	displayStudentList()

	findHighestCGPAStudent()
	avgCGPA()
}

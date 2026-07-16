package student

type Student struct {
	name       string
	age        int
	department string
	cgpa       float32
}

func registerStudent() (Student, Student) {
	student1 := Student{
		name:       "Nishak",
		age:        23,
		department: "CSE",
		cgpa:       3.19,
	}
	student2 := Student{
		name:       "Jihan",
		age:        25,
		department: "CSE",
		cgpa:       3.18,
	}
	return student1, student2
}

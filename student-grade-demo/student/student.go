package student

import "fmt"

type Student struct {
	Name  string
	Marks []int
}

func (s *Student) AddMarks(marks ...int) {
	s.Marks = append(s.Marks, marks...)
}

func (s *Student) Average() float64 {
	if len(s.Marks) == 0 {
		return 0
	}
	total := 0

	for _, mark := range s.Marks {
		total += mark
	}

	return (float64(total) / float64(len(s.Marks)))
}
func (s *Student) Display() {
	fmt.Println("Name", s.Name)
	fmt.Println("Marks", s.Marks)
	fmt.Println("Average", s.Average())
}

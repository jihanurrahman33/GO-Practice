package main

import "fmt"

func main() {
	result, err := Divide(10, 2)

	if err != nil {
		fmt.Println("Divide by zero error")
	}
	fmt.Println(result)
	result, err = Divide(10, 0)
	if err != nil {
		fmt.Println("Divide by zero error")
	}
	fmt.Println(result)
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Can't divide %d by zero", a)

	}

	return a / b, nil
}

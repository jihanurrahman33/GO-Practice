package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5}
	sum := 0
	max := 0

	calculateSum(&number, &sum)
	fmt.Println(sum)

	findMaxNum(&number, &max)
	fmt.Println("Max num is:", max)
	oddCount := 0
	evenCount := 0
	countOddEven(&number, &oddCount, &evenCount)

	fmt.Printf("total odd number is %d and even number is %d in numbers list", oddCount, evenCount)
	fmt.Println()

	secondLargest := findSecondLargestNum(&number)

	fmt.Println("Second Largest Number is", secondLargest)

	reverseASlice(&number)
	fmt.Println(number)

	rmvVal := 4
	removeElFrmSlice(&number, &rmvVal)
	fmt.Println(number)

}

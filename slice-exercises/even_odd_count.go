package main

func countOddEven(num *[]int, oddCount *int, evenCount *int) {

	for _, val := range *num {
		if val%2 == 0 {
			*evenCount++
		} else {
			*oddCount++
		}
	}

}

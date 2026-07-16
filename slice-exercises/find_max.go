package main

func findMaxNum(num *[]int, max *int) {
	for _, val := range *num {
		if val > *max {
			*max = val
		}
	}
}

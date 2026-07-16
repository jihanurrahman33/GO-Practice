package main

func reverseASlice(num *[]int) {
	left := 0
	right := len(*num) - 1
	for left < right {
		(*num)[left], (*num)[right] = (*num)[right], (*num)[left]
		left++
		right--
	}
}

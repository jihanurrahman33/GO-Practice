package main

func findSecondLargestNum(nums *[]int) int {
	largest := (*nums)[0]
	secondLargest := (*nums)[0]
	for _, num := range *nums {
		if num > largest {
			secondLargest = largest
			largest = num
		} else if num > secondLargest && num != largest {
			secondLargest = num
		}
	}

	return secondLargest
}

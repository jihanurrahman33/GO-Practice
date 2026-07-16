package main

func calculateSum(number *[]int, sum *int) {

	for _, val := range *number {
		*sum += val
	}

}

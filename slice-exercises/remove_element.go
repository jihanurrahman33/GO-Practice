package main

func removeElFrmSlice(nums *[]int, rmvVal *int) {
	rmvValIdx := 0

	for idx, val := range *nums {
		if val == *rmvVal {
			rmvValIdx = idx
		}
	}

	*nums = append((*nums)[:rmvValIdx], (*nums)[rmvValIdx+1:]...)
}

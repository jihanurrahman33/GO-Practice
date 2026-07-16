package main

import "fmt"

func printLoops(number []int) {
	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])
	}

	for idx, val := range number {
		fmt.Println(idx, val)
	}

	for _, val := range number {
		fmt.Println(val)
	}
}

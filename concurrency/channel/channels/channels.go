package main

import "fmt"

// example of unbuffered channel for sending and receiving value

func sum(done chan bool, result chan int, a, b int) {

	defer func() { done <- true }()
	fmt.Println("Calculating...")
	result <- a + b

}
func main() {
	done := make(chan bool)
	result := make(chan int)

	go sum(done, result, 10, 20)

	fmt.Println(<-result)

	<-done

}

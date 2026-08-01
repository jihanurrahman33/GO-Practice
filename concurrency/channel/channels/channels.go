package main

import "fmt"

// example of unbuffered channel for sending and receiving value
func sum(done chan bool, result chan int, a, b int) {

	defer func() { done <- true }()
	fmt.Println("Calculating...")
	result <- a + b

}

func main() {

	// example of unbuffered channel for sending and receiving value
	done := make(chan bool)
	result := make(chan int)

	go sum(done, result, 10, 20)

	fmt.Println(<-result)

	<-done

	//example of buffered channel
	emailChan := make(chan string, 100)

	emailChan <- "nishak@example.com"
	emailChan <- "jihan@example.com"
	emailChan <- "Nishakjr@example.com"

	fmt.Println("Data receiving from channel without blocking/deadlock", <-emailChan)
	fmt.Println("Data receiving from channel without blocking/deadlock", <-emailChan)
	fmt.Println("Data receiving from channel without blocking/deadlock", <-emailChan)

}

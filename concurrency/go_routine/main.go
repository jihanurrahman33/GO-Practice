package main

import (
	"fmt"
	"time"
)

func sayHello(message string, delay time.Duration) {

	time.Sleep(delay)
	fmt.Println("sayHello", message)
}

func main() {

	fmt.Println("Hello from Main() Goroutine")

	go sayHello("Hello World!", time.Second)

	fmt.Println("Last message from main() goroutine")

	time.Sleep(2 * time.Second)

}

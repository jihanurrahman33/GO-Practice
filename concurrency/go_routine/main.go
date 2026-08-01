package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(message string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println("sayHello", message)
}

func sayHello2(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("sayHello", message)
}

func main() {

	var wg sync.WaitGroup
	startTime := time.Now()
	fmt.Println("Hello from Main() Goroutine")
	wg.Add(6)
	go sayHello("Hello World! 1", time.Second, &wg)
	go sayHello("Hello World! 2", time.Second, &wg)
	go sayHello("Hello World! from 2 seconds", 2*time.Second, &wg)
	go sayHello("Hello World! from 3 seconds", 3*time.Second, &wg)
	go sayHello("Hello World! from 4 seconds", 3*time.Second, &wg)
	go sayHello("Hello World! from 5 seconds", 3*time.Second, &wg)

	wg.Wait()
	fmt.Println("Last message from main() goroutine")
	fmt.Println("Total time taken by sayHello():", time.Since(startTime))

	startTime = time.Now()

	sayHello2("Hello World! 1", time.Second)
	sayHello2("Hello World! 2", time.Second)
	sayHello2("Hello World! from 2 seconds", 2*time.Second)
	sayHello2("Hello World! from 3 seconds", 3*time.Second)
	sayHello2("Hello World! from 4 seconds", 3*time.Second)
	sayHello2("Hello World! from 5 seconds", 3*time.Second)

	fmt.Println("Total time taken by sayHello2():", time.Since(startTime))
}

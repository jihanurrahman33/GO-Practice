package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//more channel examples
	var wg sync.WaitGroup
	wg.Add(5)
	resultChan := make(chan string, 5)

	go processSomething("Hello World 1", &wg, resultChan)
	go processSomething("Hello World 2", &wg, resultChan)
	go processSomething("Hello World 3", &wg, resultChan)
	go processSomething("Hello World 4", &wg, resultChan)
	go processSomething("Hello World 5", &wg, resultChan)

	wg.Wait()
	close(resultChan)
	for result := range resultChan {
		fmt.Println(result)
	}
}

func processSomething(msg string, wg *sync.WaitGroup, resultChan chan string) {

	defer wg.Done()

	time.Sleep(time.Second)
	fmt.Println("Proccessing Something...")

	resultChan <- msg

}

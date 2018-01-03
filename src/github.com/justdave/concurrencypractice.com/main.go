package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup
var data chan string

func main() {
	fmt.Println("Starting the application...")
	data = make(chan string)

	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go worker()
	}

	for i := 0; i < 10; i++ {
		data <- ("Testing " + strconv.Itoa(i))
	}

	close(data)

	waitGroup.Wait()
}

func worker() {
	fmt.Println("Goroutine worker is now starting...")
	defer func() {
		fmt.Println("Destroying the worker...")
		waitGroup.Done()
	}()
	for {
		value, ok := <-data
		if !ok {
			fmt.Println("The channel is closed!")
			break
		}
		fmt.Println(value)
		time.Sleep(time.Second * 1)
	}
}

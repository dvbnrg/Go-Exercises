package main

import (
	"fmt"
	"time"
)

// SendValue bla bla bla
func SendValue(c chan string) {
	fmt.Println("Executing Goroutine")
	time.Sleep(1 * time.Second)
	c <- "Hello World"
	fmt.Println("Finished Executing Goroutine")
}

func main() {
	fmt.Println("Channels Bruh")

	values := make(chan string, 2)
	defer close(values)

	go SendValue(values)
	go SendValue(values)

	value := <-values
	fmt.Println(value)
}

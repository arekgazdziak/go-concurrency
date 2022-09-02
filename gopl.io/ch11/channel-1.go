package main

import (
	"fmt"
)

func main() {
	intStream := make(chan int)

	go func() {
		value := 69
		fmt.Printf("Sending to channel... %d\n", value)
		intStream <- value
	}()

	value, ok := <-intStream
	fmt.Printf("[Status: %v] Received from channel: %v\n", ok, value)
}

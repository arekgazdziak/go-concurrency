package main

import (
	"fmt"
)


func main() {
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		for i := 0; i <= 5; i++ {
			intStream <- i
		}
	}()
	
	for value := range intStream {
		fmt.Println(value)
	}
}

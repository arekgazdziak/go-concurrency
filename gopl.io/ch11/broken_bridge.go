package main

import (
	"fmt"
	"sync"
)

func main() {
	i := 5
	stream := make(chan int, 1)
	stream <- i
	close(stream)
	value := <-stream
	value = <-stream

	fmt.Print(value)

	stream2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		stream2 <- i
		close(stream2)
	}()

	fmt.Println(<-stream2)
	wg.Wait()
	fmt.Println(<-stream2)
}

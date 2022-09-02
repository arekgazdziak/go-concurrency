package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	generateValues := func(done <-chan interface{}) <-chan int {
		values := make(chan int)
		go func() {
			defer fmt.Println("All values generated")
			defer close(values)
			for {
				select {
				case <-done:
					fmt.Println("DONE")
					return
				case values <- rand.Int():
					fmt.Println("Generated")
				}

			}

		}()
		return values
	}

	done := make(chan interface{})
	values := generateValues(done)

	for i := 0; i < 3; i++ {
		fmt.Printf("Received value: %v\n", <-values)
	}
	fmt.Println("Closing the channel")
	close(done)
	time.Sleep(5 * time.Second)
}

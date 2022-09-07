package main

import "fmt"

func main() {
	asyncTask := func() <-chan int {
		finished := make(chan int)
		go func() {
			defer close(finished)
			fmt.Println("Hello")
			finished <- 1
		}()
		return finished
	}

	awaitCh1 := asyncTask()
	awaitCh2 := asyncTask()

	select {
	// receiver can test whether a channel has been closed by assigning a second parameter to the receive expression: https://go.dev/tour/concurrency/4
	case i, ok := <-awaitCh1:
		if !ok {
			return
		}
		fmt.Printf("First done: %v\n", i)
	case i, ok := <-awaitCh2:
		if !ok {
			return
		}
		fmt.Printf("Second done: %v\n", i)
	}

	select {
	case i, ok := <-awaitCh1:
		if !ok {
			return
		}
		fmt.Printf("First done: %v\n", i)
	case i, ok := <-awaitCh2:
		if !ok {
			return
		}
		fmt.Printf("Second done: %v\n", i)
	}
}

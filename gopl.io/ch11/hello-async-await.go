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
	case i := <-awaitCh1:
		if i == 0 {
			return
		}
		fmt.Printf("First done: %v\n", i)
	case i := <-awaitCh2:
		if i == 0 {
			return
		}
		fmt.Printf("Second done: %v\n", i)
	}

	select {
	case i := <-awaitCh1:
		if i == 0 {
			return
		}
		fmt.Printf("First done: %v\n", i)
	case i := <-awaitCh2:
		if i == 0 {
			return
		}
		fmt.Printf("Second done: %v\n", i)
	}
}

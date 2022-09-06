package main

import "fmt"

func main() {
	finished := make(chan int)
	go func() {
		defer close(finished)
		fmt.Println("Hello")
		finished <- 1
	}()

	i := <-finished
	fmt.Printf("It's %v, I am done\n", i)
}

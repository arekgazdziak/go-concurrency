package main

import (
	"fmt"
	"time"

	"gopl.io/ch11/functional"
)

func main() {
	signal := func(delay time.Duration) <-chan interface{} {
		done := make(chan interface{})
		go func() {
			defer close(done)
			fmt.Printf("Expected delay %v\n", delay)
			time.Sleep(delay)
			fmt.Printf("Executed after delay %v\n", delay)
		}()
		return done
	}

	<-functional.Or(
		signal(10*time.Second),
		signal(15*time.Second),
		signal(30*time.Second),
	)

}

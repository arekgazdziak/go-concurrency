package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"gopl.io/ch11/functional"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	randomizeStream := func(done <-chan interface{}) <-chan interface{} {
		randStream := make(chan interface{})

		go func() {
			defer close(randStream)
			for {
				time.Sleep(1 * time.Second)
				select {
				case <-done:
					return
				case randStream <- rand.Int():
				}
			}
		}()
		return randStream
	}

	start := time.Now()

	num := runtime.NumCPU()
	randomizers := make([]<-chan interface{}, num)
	for i := 0; i < num; i++ {
		randomizers[i] = randomizeStream(done)
	}

	for value := range functional.Take(done, functional.Merge(done, randomizers...), 5) {
		fmt.Printf("%v\n", value)
	}

	fmt.Printf("\nElapsed: %s\n", time.Since(start))

}

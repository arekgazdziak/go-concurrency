package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	var count int
	var reps int = 5

	inc := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Counter value: %d\n", count)
	}

	dec := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Counter value: %d\n", count)
	}

	for i := 0; i < reps; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inc()
		}()
	}

	for i := 0; i < reps; i++ {
		wg.Add(1)
		go func() {
			wg.Done()
			dec()
		}()
	}

	wg.Wait()
}

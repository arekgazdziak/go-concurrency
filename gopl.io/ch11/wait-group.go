package main

import (
	"sync"
	"fmt"
)

func main() {

	hello := func(id int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Printf("Hello %d\n", id)
	}

	var wg sync.WaitGroup
	const numOfCalls int = 5
	wg.Add(5)
	for i := 0; i < numOfCalls; i++ {
		go hello(i+1, &wg)
	}

	wg.Wait()
}

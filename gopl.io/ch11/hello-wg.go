package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go sayHello(&wg)
	wg.Wait()
}

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello")
}

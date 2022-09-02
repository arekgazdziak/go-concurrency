package main

import (
	"fmt"
	"time"

	"gopl.io/ch11/functional"
)

func main() {
	done := make(chan interface{})
	defer close(done)
	start := time.Now()
	stage1 := functional.Sleep(done, 1*time.Second, functional.Take(done, functional.Repeat(done, 5), 3))
	//	buffer := functional.Buffer(done, 2, stage1)
	stage2 := functional.Sleep(done, 4*time.Second, stage1)

	counter := 0

	for value := range stage2 {
		counter++
		fmt.Println(value)
	}

	fmt.Println("Elapsed time: ", time.Since(start))
	fmt.Println("Counter: ", counter)
}

package main

import (
	"fmt"
	"math/rand"
	"time"

	"gopl.io/ch11/functional"
)

func main() {
	done := make(chan interface{})
	var message string
	go func() {
		for value := range functional.ToString(done, functional.Take(done, functional.Repeat(done, "Today", "is", "the", "second", "day", "of", "go", "pipeline"), 5)) {
			message += fmt.Sprintf(" %v", value)
			fmt.Println(message)
		}
	}()

	fn := func() interface{} { return rand.Int() }
	go func() {
		for value := range functional.Take(done, functional.RepeatFn(done, fn), 50) {
			fmt.Printf("%v\n", value)
		}
	}()

	time.Sleep(10000 * time.Microsecond)
	fmt.Println(message)
	close(done)
}

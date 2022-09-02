package main

import (
	"fmt"
	"time"
)

func main() {
	analyseSync := func(done <-chan interface{}, values <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer close(terminated)
			for {
				select {
				case <-done:
					return
				case value := <-values:
					fmt.Printf("Received value: %v\n", value)
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	syncData := make(chan string)
	terminated := analyseSync(done, syncData)

	syncData <- "My values"

	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	syncData <- "My better value"

	<-terminated
	fmt.Println("Thanks")
}

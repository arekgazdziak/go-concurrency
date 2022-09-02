package main

import (
	"fmt"
	"time"

	"gopl.io/ch11/heartbeats"
)

func main() {
	done := make(chan interface{})
	defer close(done)
	const interval = 2 * time.Second

	heartbeats, result := heartbeats.Tick(done, interval)

	for {
		select {
		case _, ok := <-heartbeats:
			if ok == false {
				fmt.Println("Error occured")
				return
			}
			fmt.Println("Pulse")
		case item, ok := <-result:
			if ok == false {
				return
			}
			fmt.Println("Received value: ", item.Format(time.RFC850))
		case <-time.After(interval):
			return
		}

	}

}

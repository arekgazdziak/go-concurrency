package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	done := make(chan struct{})
	ch1 := make(chan string)
	ch2 := make(chan string)
	interval := 10 * time.Millisecond
	timeout := 10 * time.Second
	pulseTimeout := 2 * interval

	go goBackgroundJob(1, ch1, done, interval)
	go goBackgroundJob(2, ch2, done, interval)

	go func() {
		<-time.After(timeout)
		fmt.Println("TASK TIEMOUT")
		close(done)
	}()

loop:
	for {
		timeoutChannel := time.After(pulseTimeout)
		for {
			select {
			case msg1 := <-ch1:
				fmt.Println(msg1)
			case msg2 := <-ch2:
				fmt.Println(msg2)
				continue loop
			case <-timeoutChannel:
				fmt.Println("pulse timeout")
				continue loop
			case <-done:
				return
			}
		}
	}
}

func goBackgroundJob(id int, ch chan string, done <-chan struct{}, pulseInterval time.Duration) {
	pulse := time.Tick(pulseInterval)
	value := 0
	for {
		select {
		case t := <-pulse:
			value++
			fmt.Println(t)
			ch <- strconv.Itoa(value) + " *** " + strconv.Itoa(id)
		case <-done:
			return
		}
	}
}

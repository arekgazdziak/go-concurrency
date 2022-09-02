package functional

import "fmt"

func Bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)

		for {
			stream := make(<-chan interface{})
			select {
			case <-done:
				return
			case maybeStream, ok := <-chanStream:
				if ok == false {
					fmt.Println("ClosedStream")
					return
				}
				stream = maybeStream
			}
			for val := range stream {
				select {
				case <-done:
					return
				case valStream <- val:
				}
			}
		}
	}()
	return valStream
}

package functional

import (
	"fmt"
	"sync"
	"time"
)

func Repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

func Take(done <-chan interface{}, values <-chan interface{}, number int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < number; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-values:
			}
		}
	}()
	return takeStream
}

func ToString(done <-chan interface{}, values <-chan interface{}) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for v := range values {
			select {
			case <-done:
				return
			case stringStream <- v.(string):
			}
		}
	}()
	return stringStream
}

func RepeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valuesStream := make(chan interface{})
	go func() {
		defer close(valuesStream)
		for {
			select {
			case <-done:
				return
			case valuesStream <- fn():
			}
		}
	}()
	return valuesStream
}

func Merge(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
	mergedStream := make(chan interface{})
	var wg sync.WaitGroup

	readSingleStream := func(singleStream <-chan interface{}) {
		defer wg.Done()

		for item := range singleStream {
			select {
			case <-done:
				return
			case mergedStream <- item:
			}
		}
	}

	wg.Add(len(channels))
	for _, channel := range channels {
		go readSingleStream(channel)
	}

	go func() {
		wg.Wait()
		close(mergedStream)
	}()

	return mergedStream
}

func Sleep(done <-chan interface{}, duration time.Duration, values <-chan interface{}) <-chan interface{} {
	streamValues := make(chan interface{})
	go func() {
		defer close(streamValues)
		for value := range values {
			select {
			case <-done:
				return
			case <-time.After(duration):
				fmt.Printf("Received value: %v\n", value)
				streamValues <- value
			}
		}
	}()

	return streamValues
}

func Buffer(done <-chan interface{}, size int, values <-chan interface{}) <-chan interface{} {
	buffer := make(chan interface{}, 2)
	go func() {
		defer close(buffer)
		for value := range values {
			select {
			case <-done:
				return
			case buffer <- value:
				fmt.Println("Buffering...")
			}

		}
	}()
	return buffer
}

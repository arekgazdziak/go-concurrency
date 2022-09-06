package heartbeats

import "time"

func Tick(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
	heartbeats := make(chan interface{})
	result := make(chan time.Time)

	go func() {
		defer close(heartbeats)
		defer close(result)

		pulseStream := time.Tick(pulseInterval / 2)
		timeStream := time.Tick(pulseInterval)

		handlePuls := func() {
			select {
			case <-done:
				return
			case heartbeats <- struct{}{}:
			default:
			}
		}

		handleWork := func(time time.Time) {
			for {
				select {
				case <-done:
					return
				case <-pulseStream:
					handlePuls()
				case result <- time:
					return
				}
			}
		}

		for {
			select {
			case <-done:
				return
			case <-pulseStream:
				handlePuls()
			case time := <-timeStream:
				handleWork(time)
			}
		}
	}()

	return heartbeats, result
}

func TickTest(done <-chan interface{}, pulseInterval time.Duration, numbers ...int) (<-chan interface{}, <-chan int) {
	heartbeats := make(chan interface{})
	result := make(chan int)

	go func() {
		defer close(heartbeats)
		defer close(result)

		pulseStream := time.Tick(pulseInterval)
	numLoop:
		for _, n := range numbers {

			for {
				select {
				case <-done:
					return
				case <-pulseStream:
					select {
					case heartbeats <- struct{}{}:
					default:
					}
				case result <- n:
					continue numLoop
				}
			}
		}

	}()

	return heartbeats, result
}

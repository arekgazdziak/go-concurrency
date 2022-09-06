package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch11/functional"
)

type superviseeFn func(
	done <-chan interface{},
	pulseInterval time.Duration,
) (heartbeat <-chan interface{})

func main() {
	newSupervisor := func(
		timeout time.Duration,
		childGoroutine superviseeFn,
	) superviseeFn {
		return func(
			done <-chan interface{},
			pulseInterval time.Duration,
		) <-chan interface{} {
			heartbeat := make(chan interface{})

			go func() {
				defer close(heartbeat)

				var childHeartbeat <-chan interface{}
				var childDone chan interface{}

				startChild := func() {
					childDone = make(chan interface{})
					childHeartbeat = childGoroutine(functional.Or(done, childDone), pulseInterval)
				}

				startChild()
				pulse := time.Tick(pulseInterval)

			monitorLoop:
				for {
					timeoutSignal := time.After(timeout)
					for {
						select {
						case <-pulse:
							select {
							case heartbeat <- struct{}{}:
							default:
								log.Println("supervisor: hearbeat default")
							}
						case <-childHeartbeat:
							log.Println("child: hearbeat received")
							continue monitorLoop
						case <-timeoutSignal:
							log.Println("supervisor: child unhealthy, restarting")
							close(childDone)
							startChild()
							continue monitorLoop
						case <-done:
							return
						}
					}
				}
			}()
			return heartbeat
		}
	}

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	childWork2 := func(done <-chan interface{}, values ...int) (superviseeFn, <-chan interface{}) {
		intChanStream := make(chan (<-chan interface{}))
		intStream := functional.Bridge(done, intChanStream)

		doWork := func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} {
			intStream := make(chan interface{})
			heartBeat := make(chan interface{})

			go func() {
				defer close(intStream)
				select {
				case <-done:
					return
				case intChanStream <- intStream:
				}

				pulse := time.Tick(pulseInterval)

				for {
				valueLoop:
					for _, intVal := range values {
						if intVal < 0 {
							log.Printf("negative value: %v\n", intVal)
							return
						}
						for {
							select {
							case intStream <- intVal:
								continue valueLoop
							case <-done:
								return
							case <-pulse:
								select {
								case heartBeat <- struct{}{}:
								default:
								}

							}
						}
					}
				}

			}()
			return heartBeat
		}

		return doWork, intStream
	}

	done := make(chan interface{})

	task, intStream := childWork2(done, 1, 2, 3, 7, 4, 5)

	taskWithSupervisor := newSupervisor(20*time.Millisecond, task)

	taskWithSupervisor(done, 10*time.Millisecond)

	time.AfterFunc(10*time.Second, func() {
		log.Println("main, halting supervisor and its childs")
		close(done)
	})

	for value := range intStream {
		_ = value
	}

	fmt.Println("Done")
}

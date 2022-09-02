package heartbeats

import (
	"testing"
	"time"
)

func Test_Ticker(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 4, 5}
	const timeout = 2 * time.Second

	heartbeats, result := TickTest(done, timeout/2, intSlice...)

	<-heartbeats

	i := 0
	for {
		select {
		case r, ok := <-result:
			if ok == false {
				return
			} else if expected := intSlice[i]; r != expected {
				t.Errorf(
					"index %v: expected: %v, but received: %v",
					i,
					expected,
					r,
				)

			}
			i++
		case <-heartbeats:
		case <-time.After(timeout):
			t.Fatal("timeout")
		}
	}

}

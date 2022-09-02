package functional

func Or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	done := make(chan interface{})

	go func() {
		defer close(done)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		case 3:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-Or(append(channels[3:], done)...):
			}
		}
	}()
	return done
}

func OrDone(done, c <-chan interface{}) <-chan interface{} {
	values := make(chan interface{})
	go func() {
		defer close(values)

		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case <-done:
					return
				case values <- v:
				}

			}

		}

	}()
	return values
}

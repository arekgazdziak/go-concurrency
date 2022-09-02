package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

}

func goOne(ch chan string) {
	for i := 0; i < 2; i++ {
		fmt.Println("ONE: Ready to write")
		ch <- "From goOne goroutine"
	}
}

func goTwo(ch chan string) {
	for i := 0; i < 2; i++ {
		fmt.Println("TWO: Ready to write")
		ch <- "From goTwo goroutine"
	}
}

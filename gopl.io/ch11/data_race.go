package main

import "fmt"

func main() {
	var data int

	go func() {
		data++
	}()

	fmt.Println(data)

}

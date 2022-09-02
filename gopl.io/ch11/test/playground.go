package main

import (
	"fmt"
)

func main() {
	set := make(map[int]struct{})

	// Add some values
	set[1] = struct{}{}
	set[5] = struct{}{}

	if _, ok := set[1]; ok {
		fmt.Println("It's there")
	}
}

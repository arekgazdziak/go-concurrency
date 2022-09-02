package main

import (
	"fmt"
	"gopl.io/ch11/test/format"
	"os"
	"time"
)

func main() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	var newCar = struct {
		Make    string `json:"make"`
		Model   string `json:"model"`
		Mileage int    `json:"mileage"`
	}{}
	fmt.Println(format.Any(x))                  // "1"
	fmt.Println(format.Any(d))                  // "1"
	fmt.Println(format.Any([]int64{x}))         // "[]int64 0x8202b87b0"
	fmt.Println(format.Any([]time.Duration{d})) // "[]time.Duration 0x8202b87e0"
	fmt.Println(format.Any([]time.Duration{d})) // "[]time.Duration 0x8202b87e0"
	fmt.Println(format.Any(newCar))             //
	format.Display("os.Stderr", os.Stderr)

}

package main

import (
	"fmt"
	"gopl.io/ch11/sexpr"
)

func main() {
	result,_ := sexpr.Marshal("arek")
	fmt.Println(string(result))
}:wq

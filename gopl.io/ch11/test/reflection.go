package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
        fmt.Printf("%T is the type of %v\n", 3, 3)
        v := reflect.ValueOf(3) // a reflect.Value
        fmt.Printf("%v\n", v)
        fmt.Println(v.String())
}


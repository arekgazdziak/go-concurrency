package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	b := reflect.ValueOf(x)
	y := reflect.ValueOf(&x)
	fmt.Println(x)
	fmt.Println(b)
	fmt.Println(&x)
	fmt.Println(&b)
	fmt.Println(reflect.TypeOf(y))
	fmt.Println(y)

	var px *int = y.Elem().Addr().Interface().(*int)
	fmt.Println(reflect.TypeOf(px))
	fmt.Println(px)
	fmt.Println(x)
	
	fmt.Println("*******")
	var c *Camel = &Camel{Name: "arek", Surname: "gazdziak"}
	fmt.Printf("%v\n", c)
	w := MakeWalker(c)
	fmt.Printf("%v\n", w)
	fmt.Printf("%v\n", reflect.TypeOf(w))
	fmt.Printf("%v\n", reflect.ValueOf(w))
	
}

type Walker interface {
    Walk(miles int)
}

type Camel struct {
    Name string
    Surname string
}

func (c *Camel) Walk(miles int) {
     fmt.Printf("%s is walking %v miles\n", c.Name, miles)
}

func MakeWalker(c *Camel) Walker {
	fmt.Printf("%v\n", c)
	return c
}

package main

import (
	"fmt"
)


func main() {
	name := "arek"
	name2 := "radek"

	print_my_value(&name)

	swap(&name, &name2)

	
	fmt.Printf("name address: %v\n", &name)
	fmt.Printf("name2 address: %v\n", &name2)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("name2: %v\n", name2)
}

func print_my_value(name *string) {
	fmt.Println(name)
}

func swap(first *string, second *string) {
	temp := *second
	*second = *first
	*first = temp
	
}

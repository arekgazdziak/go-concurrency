package main

import (
    "fmt"
)

func main() {
    letters1 := make([]rune, 0, 10)
    letters2 := make([]rune, 9, 10)
    letters3 := make([]rune, 10)

    runes := [][]rune{letters1, letters2, letters3 }

    for _, rune := range runes {
        fmt.Println("Len: %v", len(rune))
        fmt.Println("Cap: %v", cap(rune))
    }
}


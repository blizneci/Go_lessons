package main

import (
    "fmt"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }

    pow1 := make([]int, 10)
    for i := range pow1 {
        pow1[i] = 1 << uint(i) // == 2**i
    }
    for _, value := range pow1 {
        fmt.Printf("%d\n", value)
    }
}

// 16 Range
// The 'range' form of the 'for' loop iterates over a slice or a map.
// When ranging over a slice, two values are returned for each iteration.
// The first is the index, and the second is a copy of the element at that index.
//
// 17 Range continued.
// You can skip the index or value by assigning to '_'.
// for i, _ := range pow
// for _, value := range pow
//
// If you only want the index, you can omit the second variable.
// for i := range pow


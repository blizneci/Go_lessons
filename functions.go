package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func add_1(x, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    fmt.Println(add(42, 13))
    fmt.Println(add_1(100, 11))
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}

// A function cat take zero or more args.
// The type comes after the var name.
// When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last.
// A func cat return any number of results.
// The 'swap' func returns two strings.

package main

import (
    "fmt"
)


func main() {
    fmt.Println("counting")

    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("done")
}


// Deferred func calls are pushed onto a stack.
// When a func returns, its deferred calls are executed in last-in-first-out order.


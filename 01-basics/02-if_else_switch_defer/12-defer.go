package main

import (
    "fmt"
)


func main() {
    defer fmt.Print("world")

    fmt.Print("hello ")
}


// a 'defer' stmt defers the execution of a func until the surrounding func
// returns.
// The deferred call's arguments are evaluated immediately, but the func call
// is not executed until the surrounding func returns.


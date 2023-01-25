package main

import (
    "fmt"
    "runtime"
)


func main() {
    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("Linux")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s.\n", os)
    }
}


// A 'switch' stmt is a shorter way to write a sequence of 'if-else' stmts.
// It runs the first case whose value is equal to the condition expr.
// Go's 'switch' is like the one in C, C++, Java, JavaSctipt, and PHP, except
// that Go onlly runs the selected case, not all the cases that follow.
// In effect, the 'break' stmt that is needed at the end of each case in those
// languages is provided automatically in Go.
// Another important difference is that Go's switch cases need not be constants,
// and the values involved need not be integers.


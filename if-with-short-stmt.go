package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}

func main() {
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
}


// Like 'for', the 'if' stmt can start with short stmt to execute
// before the condition.
// Vars declared by the stmt are only in scope until the end of the 'if'.


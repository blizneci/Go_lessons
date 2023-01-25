package main

import (
    "fmt"
    "math"
)


func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func main() {
    fmt.Println(sqrt(2), sqrt(-4))
}


// Go's 'if' stmts are like its 'for' loops; the expr need not be
// surrounded by paretheses '()' but the braces '{}' are required.


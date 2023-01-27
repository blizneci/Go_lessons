package main

import (
    "fmt"
)

// fibonacci is a function that returns
// a function that returns an int.

func fibo() func() float64 {
    first := -1.0
    second := 1.0
    return func() float64 {
        first, second = second, first + second
        return second
    }
}

func main() {
    f := fibo()
    for i := 1; i <= 10; i++ {
        fmt.Printf("%d : %20.f\n", i, f())
    }
}

//


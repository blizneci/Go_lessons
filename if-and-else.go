package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("%g >= %g\n", v, lim)
    }
    // can't use v here, though
    return lim
}

func main() {
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
}


// Vars declared inside an 'if' short stmt are also available inside any
// of the 'else' blocks.
// Both calls to 'pow' return their results before the call to 'fmt.Println'
// in 'main' begins.
// Так как сначала выполняются функции 'pow', то сначала будет выведена
// строка '%g >= %g'  с соответствующими заполнителями
// а затем результаты выполнения функций.

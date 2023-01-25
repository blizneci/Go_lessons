package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    i := 1
    var e float64 = 0.00000000001
    /*
    for i := 0; i <= 10; i++ {
        z -= (z*z - x) / (2*z)
        fmt.Println(z)
    }
    */
    for math.Abs((z*z - x) / (2*z)) >= e {
        fmt.Printf("%v : %v\n", i, z)
        z -= (z*z - x) / (2*z)
        i += 1
    }
    return z
}

func main() {
    var a float64 = 256
    fmt.Println(Sqrt(a))
    fmt.Println("From the math.Sqrt is:", math.Sqrt(a))
}


//


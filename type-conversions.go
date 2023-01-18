package main

import (
    "fmt"
    "math"
)


func main() {
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x*x + y*y))
    var z uint = uint(f)
    fmt.Println(x, y, z)
}


// The expression T(v) converts the value 'v' to type 'T'.
// var i int = 42
// var f float64 = float64(i)
// var u uint = uint(f)
// or, put more simply:
// i := 42
// f := float64(i)
// u := uint(f)

// In Go assignment between items of different type requires an explicit
// conversion.
// If you remove conversions in the example you will see an error:
// ./prog.go:11:15: cannot use f (variable of type float64) as type uint in variable declaration


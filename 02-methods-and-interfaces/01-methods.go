package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

type MyFloat float64

// 01
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// 02
func Abs2(v Vertex) float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// 03
func (f MyFloat) Abs3() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

func main() {
    fmt.Println("01 Methods")
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
    fmt.Println("02 Methods are functions")
    v1 := Vertex{3, 4}
    fmt.Println(Abs2(v1))
    fmt.Println("03 Methods continued")
    f := MyFloat(-math.Sqrt(2))
    fmt.Println(f.Abs3())

}

// 01 Methods
// Go does not have classes. However, you can define methods on types.
// 
// A method is a function with a special 'receiver' argument.
// 
// The receiver appears in its own argument list between the func keyword 
// and the method name.
// In this example, the 'Abs' method has a receiver of type 'Vertex' named 'v'.
//
// 02 Methods are functions
// Remember: a method is just a function with a receiver argument.
// 
// Here's Abs written as a regular function with no change in functionality.
//
// 03 Methods continued
// Methods continued
// You can declare a method on non-struct types, too.
// 
// In this example we see a numeric type 'MyFloat' with an 'Abs' method.
// 
// You can only declare a method with a receiver whose type is defined in 
// the same package as the method. You cannot declare a method with a 
// receiver whose type is defined in another package (which includes the 
// built-in types such as 'int').


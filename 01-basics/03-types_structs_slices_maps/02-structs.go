package main

import (
    "fmt"
)

type Vertex struct { // this is a struct
    X int // a field 
    Y int // a field 
}
// typ3 Vertex struct { X, Y int } is also correct 

func main() {
    fmt.Println(Vertex{1, 2})

    v := Vertex{2, 3}
    v.X = 4
    fmt.Println(v.X)

    v1 := Vertex{3, 4}
    p := &v // a pointer to a struct
    p.X = 1e9 // access to the struct field 
    fmt.Println(v1)

    var (
        v2 = Vertex{1, 2} // has type Vertex 
        v3 = Vertex{X: 1} // Y:0 is implicit 
        v4 = Vertex{}     // X:0 and Y:0
        p1 = &Vertex{1, 2} // has type *Vertex
    )
    fmt.Println(v2, p1, v3, v4)
}


// Structs.
// A 'struct' is a collection of fields.

// Struct fields.
// Struct fields are accessed using a dot.

// Pointers to structs.
// Struct fields can be accessed through a struct pointer.
// To access the field 'X' of a struct when we have the struct pointer 'p' we 
// could write '(*p).X'. However, that notation is cumbersome, so the language 
// permits us instead to write just 'p.X', without the explicit dereference.

// Struct Literals
// A struct literal denotes a newly allocated struct value by listing the 
// values of its fields.
// 
// You can list just a subset of fields by using the Name: syntax. (And the 
// order of named fields is irrelevant.)
// 
// The special prefix '&'returns a pointer to the struct value.

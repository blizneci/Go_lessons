package main

import (
    "fmt"
)

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

var m1 = map[string]Vertex{
    "Bell Labs": Vertex{
        40.68433, -74.39967,
    },
    "Google": Vertex{
        37.42202, -122.08408,
    },
}

var m2 = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google": {37.42202, -122.08408},
}

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }
    fmt.Println(m["Bell Labs"])

    fmt.Println(m1)

    fmt.Println(m2)
}

// 19 Maps
// A map maps keys to values.
// 
// The zero value of a map is nil. A nil map has no keys, 
// nor can keys be added.
// 
// The make function returns a map of the given type, initialized 
// and ready for use.
//
// 20 Map literals
// Map literals are like struct literals, but the keys are required.
//
// 21 Map literals continued.
// If the top-level type is just a type name, you can omit it from
// the elements of the literal.

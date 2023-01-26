package main

import (
    "golang.org/x/tour/pic"
)

func f(x, y int) uint8 {
    result := uint8(10*x + 10*y)
    return result
}

func Pic(dx, dy int) [][]uint8 {
    table := make([][]uint8, dy)
    for y := 0; y < dy; y++ {
        row := make([]uint8, dx)
        for x := 0; x < dx; x++ {
            row[x] = f(x, y)
        }
        table[y] = row
    }
    return table
}

func main() {
    pic.Show(Pic)
}

// Exericse: Slice
// Implement Pic. It should return a slice of length 'dy', each element of
// which is a slice of 'dx' 8-bit unsigned integers. When you run a program,
// it will display your picture, interpreting the integers as grayscale vals.
// You need to use a loop to allocate each '[]uint8' insede the '[][]uint8'
// use 'uint8(intValue)' to convert between types.
// Works only on the Go playground.


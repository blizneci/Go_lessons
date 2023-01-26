package main

import (
    "fmt"
)

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
    // 07 Slices
    fmt.Println("07 Slices")
    primes := [6]int{2, 3, 5, 7, 11, 13}

    var s []int = primes[1:4]
    fmt.Println(s)

    // 08 Slices are like references to arrays
    fmt.Println("08 Slices are like ref")
    names := [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }
    fmt.Println(names)
    a := names[0:2]
    c := names[1:3]
    fmt.Println(a, c)
    c[0] = "XXX"
    fmt.Println(a, c)
    fmt.Println(names)

    // 09 Slice literals
    fmt.Println("09 Slice literals")
    q := []int{2, 3, 5, 7, 11, 13}
    fmt.Println(q)
    r := []bool{true, false, true, true, false, true}
    fmt.Println(r)
    st := []struct{
        i int
        b bool
    }{
        {2, true},
        {3, false},
        {5, true},
        {7, true},
        {11, false},
        {13, true},
    }
    fmt.Println(st)

    // 10 Slice defaults
    fmt.Println("10 Slice defaults")
    s1 := []int{2, 3, 5, 7, 11, 13}
    s1 = s1[1:4]
    fmt.Println(s1)
    s1 = s1[:2]
    fmt.Println(s1)
    s1 = s1[1:]
    fmt.Println(s1)

    // 11 Slice length and capacity 
    s2 := []int{2, 3, 5, 7, 11, 13}
    printSlice(s2)
    // Slice the slice to give it zero length.
    s2 = s2[:0]
    printSlice(s2)
    // Extend its length.
    s2 = s2[:4]
    printSlice(s2)
    // Drop its first two values.
    s2 = s2[2:]
    printSlice(s2)
}


// 07 Slices
// An array has a fixed size. A slice, on the other hand, is a 
// dynamically-sized, flexible view into the elements of an array. 
// In practice, slices are much more common than arrays.
// 
// The type []T is a slice with elements of type T.
// 
// A slice is formed by specifying two indices, a low and high bound, 
// separated by a colon:
// 
// a[low : high]
// This selects a half-open range which includes the first element, 
// but excludes the last one.
// 
// The following expression creates a slice which includes elements 1 
// through 3 of 'a':
// 
// a[1:4]
//
// 09 Slices are like references to arrays
// A slice does not store any data, it just describes a section of an 
// underlying array.
// 
// Changing the elements of a slice modifies the corresponding elements of 
// its underlying array.
// 
// Other slices that share the same underlying array will see those changes.
// 
// 09 Slice literals
// A slice literal is like an array literal without the length.
// 
// This is an array literal:
// 
// [3]bool{true, true, false}
// And this creates the same array as above, then builds a slice that 
// references it:
// 
// []bool{true, true, false}
// 
// 10 Slice defaults
// When slicing, you may omit the high or low bounds to use their 
// defaults instead.
// 
// The default is zero for the low bound and the length of the slice for 
// the high bound.
// 
// For the array
// var a [10]int
// these slice expressions are equivalent:
// a[0:10]
// a[:10]
// a[0:]
// a[:]
//
// 11 Slice length and capacity
// A slice has both a length and a capacity.
// 
// The length of a slice is the number of elements it contains.
// 
// The capacity of a slice is the number of elements in the underlying array, 
// counting from the first element in the slice.
// 
// The length and capacity of a slice s can be obtained using the 
// expressions len(s) and cap(s).
// 
// You can extend a slice's length by re-slicing it, provided it has sufficient 
// capacity. Try changing one of the slice operations in the example program 
// to extend it beyond its capacity and see what happens.


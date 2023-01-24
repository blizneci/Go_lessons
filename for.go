package main

import (
    "fmt"
)


func main() {
    sum := 0
    for i:= 0; i < 10; i++ {
        sum += i
    }
    fmt.Println("Total sum is:", sum)

    sum_1 := 1
    for ; sum_1 < 1000; {
        sum_1 += sum_1
    }
    fmt.Println("Total sum_1 is:", sum_1)

    sum_2 := 50
    for sum_2 < 1000 {
        sum_2 += sum_2
    }
    fmt.Println("Total sum_2 is:", sum_2)

    sum_3 := 1
    for {
        sum_3 += sum_3
        if sum_3 > 100 {
            break
        }
        fmt.Println("Total sum_3 is:", sum_3)
    }
}

// Go has only one looping construct, the for loop
// Three components separated by ';':
// the init stmt: executed before the first iteration
// the condition expr: evaluated before every iteration
// the post stmt: executed at the end of every iter.
// The vars declared there are visible only in the scope of the 'for' stmt.
// There are no parentheses surrounding the three components of the 'for' stmt
// and the braces '{}' are always required.

// The init and post stmts are optional.

// You can drop the ';'. C's 'while' is spelled 'for' in Go.

// If you omit loop condition it loops forever, so an inf loop is
// compactly expressed.

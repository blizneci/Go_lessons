package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - 4
    return
}

func main() {
    fmt.Println(split(17))
}

// Go's return values may be named. If so, they are treated as vars at the
// top of the func.
// These names should be used to document the meaning of the return values.
// A return stmt without args returns the named return values.
// This is known as a "naked" return.
// Naked return stmt should be used only in short funcs.
// they can harm readability in longer funcs.

package main

import "fmt"

// There is a list of vars
var c, python, java bool
// 'bool' is a type. Default is 'false'.
var j, k int = 1, 2

func main() {
    var i int
    // 'int' is a type. Default is '0'.
    fmt.Println(i, c, python, java)
    var c, python, java = true, false, "no!"
    fmt.Println(j, k, c, python, java)
    // short var declarations
    php, perl, cobol := true, false, "yes!"
    m := 3
    fmt.Println(m, php, perl, cobol)
}

// Variables:
// The 'var' stmt declares a list of vars;
// as in func args lists, the type is last.
// A 'var' stmt can be at package of func level.

// Variables with initializers:
// A var declaration can include initializers, one per var.
// If an initializer is present, the type can be omited;
// the var will take the type of the initializer.

// Short var declarations
// Inside the func, the ':=' short assignment stmt can be used in place of
// a 'var' declaration with implicit type.
// Outside a func, every stmt with begins with a keyword ('var', 'func', etc)
// and so the ':=' construc is not available.

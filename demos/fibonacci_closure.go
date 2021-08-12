package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    term1 := 0
    term2 := 1
    return func() int {
        term2 = term1 + term2
        term1 = term2 - term1
        return term1
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}

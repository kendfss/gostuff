package main
 
import (
    "fmt"
    "os"
)
 
func main() {
        // the first argument i.e. program name is excluded
    argLength := len(os.Args[1:])  
    fmt.Printf("Arg length is %d\n", argLength)
 
    for i, a := range os.Args[1:] {
        fmt.Printf("Arg %d is %s\n", i+1, a) 
    }
}
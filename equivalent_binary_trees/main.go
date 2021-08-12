package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    ch <- t.Value
    Walk(t.Left, ch)
    Walk(t.Right, ch)
    //close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    for i := 0; i >= 0; i++ {
        // for {
        v1, s1 := <-ch1
        v2, s2 := <-ch2
        fmt.Println(i, v1, v2, s1, s2)
        // if v1 != v2 {
        // return false
        // }
        if (!s1 && s2) || (s1 && !s2) {
            return false
        }
        if !s1 && !s2 {
            fmt.Println("breaking", s1, s2)
            break
        }
    }
    return true
}
func main() {
    if !Same(tree.New(1), tree.New(1)) {
        panic("not Same(tree.New(1), tree.New(1))")
    }
    if Same(tree.New(1), tree.New(2)) {
        panic("Same(tree.New(1), tree.New(2))")
    }

}

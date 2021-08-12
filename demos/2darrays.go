package main

// import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy uint) [][]float32 {
    p := make([][]float32, dy)
    for i := range p {
        p[i] = make([]float32, dx)
    }
    return p    
}

func main() {
    // pic.Show(Pic)
    var i uint = 1
    for ;i < 10; i++ {
        fmt.Println(Pic(3, i))    
    }
    
    // fmt.Println(Pic(3, 4))
}


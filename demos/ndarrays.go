package main

// import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
    // return make(make([]uint8, dx), dy)
    p := make([][]uint8, dy)
    for i := range p {
        p[i] = make([]uint8, dx)
    }
    return p
}

func main() {
    // pic.Show(Pic)
    fmt.Println(Pic(1, 2))
    fmt.Println(Pic(3, 4))
}


package main


import "fmt"

// type Particle []float64 {
type Particle struct {
    Mass float64
    Position []float64
    // []float64
}
func (self Particle) Moment(i int) float64 {
    return self.Mass * self.Position[i]
}






func main() {
    e := Particle{0., []float64 {0., 0., 0.}}
    fmt.Println(e)
    fmt.Println(e[0])
}
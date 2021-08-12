package main

import (
    "fmt"
    "math"
    "math/rand"
)

func delta(value, target float64) float64 {
    return (value*value - target) / (2 * value)
}

// func Sqrt(x float64) float64 {
//     z := 1. - rand.Float64()
//     for math.Abs(z*z - x) >= math.Pow(10, -13) {
//         z -= (z*z - x) / (2 * z)
//     }
//     return z
// }

type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}
func Sqrt(x float64) (float64, error) {
    if x >= 0 {
        z := 1. - rand.Float64()
        for math.Abs(z*z - x) >= math.Pow(10, -13) {
            z -= (z*z - x) / (2 * z)
        }
        return z, nil
    } else {
        return 0, ErrNegativeSqrt(x)
    }
}

func report(value float64) {
    fmt.Println(value)
    // v, err := Sqrt(value)
    v, err := Sqrt(value)
    // if err == ErrNegativeSqrt(value) {
    if err.(ErrNegativeSqrt) == err {
        print(ErrNegativeSqrt(value))
    } else {
        fmt.Println(math.Sqrt(value))
        fmt.Println(v)
        fmt.Println(v - math.Sqrt(value))
        fmt.Println("")
    }
}

func main() {
    // x := 4.
    for i := 2.; i < 120.; i++ {
        report(i)
    }
}

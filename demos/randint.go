package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(rand.Intn(3))
	// }
	fmt.Println(rand.Int() % 3)
	fmt.Println(rand.Int() % 3)
	fmt.Println(rand.Int() % 3)
	fmt.Println(rand.Int() % 3)
	fmt.Println(rand.Int() % 3)
	fmt.Println(Choice(1, 2, 3, 4))
	var x = iota
	fmt.Println(x)
}

func Choice(options...interface{}) interface{} {
	index := rand.Int() % len(options)
	return options[index]
}
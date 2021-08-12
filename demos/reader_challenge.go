package main

import (
    "fmt"
)

type MyReader struct{}

// func (MyReader) Read(b []byte) (n int, err error) {
// 	b
// }






func main() {
	var a [][]byte
	fmt.Println(a)
	b := []byte("A")
	fmt.Println(b)
	append(a, b)
	fmt.Println(a)
	// fmt.Println([]byte("A"))
	// var b []byte
    // fmt.Println(b)
	// append(b, []byte("A"))
    // fmt.Println(b)
}
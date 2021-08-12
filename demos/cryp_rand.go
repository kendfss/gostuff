package main


import (
	"fmt"
	"io"
	"crypto/rand"
	"math/big"
)



func main() {
	b := make([]byte, 5)
	var r io.Reader
	// var i &qbig.Int
	i := (big.NewInt(2))
	fmt.Println(rand.Read(b))
	// fmt.Println(rand.Int(io.Reader, 1000))
	fmt.Println(rand.Int(r, i))
}
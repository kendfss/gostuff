package main

import (
	_"log"
)
var alphabet word
var bank word 
type word []rune 

func (self word) freq(char rune) (val int) {
	for _, elem := range self {
		if elem == char {
			val += 1
		}
	}
	return
}
func (self word) weights() (rack []float64) {
	for _, elem := range self {
		/* code */
	}
}
func (self word) choice() {}
func freq(weight float64) float64 {}

func new_word() {
	if char := alphabet.choice(); char == ' ' {
		return char
	} else {
		return char + new_word
	}
}






func main() {
	
}

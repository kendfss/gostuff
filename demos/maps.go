package main
import (
	// "golang.org/x/tour/wc"
	"fmt"
)


const lorem string = "Consequat ex quis magna occaecat culpa sunt nulla nisi in duis consequat id esse labore ut."


func split(str, sep string) []string {
	// var bank = make([]string, 1) // = []
	// var bank [][]string
	var bank = make([]string, 1)
	var j uint
	for i, v := range str {
		if (v == sep) {
			// append(bank, []string)
			append(bank, "")
			j++
			
		} else {
			bank[j] += v
		}
	}
	return bank
}



func main() {
	fmt.Println(lorem)
	fmt.Println(split(lorem, " "))
}
// selects the top N results from a list

package main

import "log"
import "crypto/rand"

func Clone(slice []int) []int {
	clone := make([]int, len(slice))
	for i, elem := range slice {
		clone[i] = elem
	}
	return clone
}
func TopN(board []int, score, index int) []int {
	if index < len(board)-1 {
		elem := board[index]
		if (elem == 0 && score != 0) || score < elem {
			board[index] = score
			if elem == 0 {
				return board
			}
			score = elem
		}
		return TopN(board, score, index+1)
	}
	return board 
}

func Range(start, stop, step int) []int {
	var rack []int
	for ; start < stop; start+=step {
		rack = append(rack, start)
	}
	return rack
}
func pickle(uo...interface{}) {}
func Swap(slice []int, m, n int) []int {
	slice[m], slice[n] = slice[n], slice[m]
	return slice
}
func IndexPair(max int) [2]int {
	pair := [2]int{}
	for pair[0] == pair[1] {
		pair[0] = Randex(max)
		pair[1] = Randex(max)
	}
	return pair
}
func Randex(length int) int {
    bites := make([]byte, 1)
    _, err := rand.Read(bites)
    if err != nil {
    	log.Fatal("Couldn't generate random index")
    }
    return int(bites[0]) % length
}
func Shuffle(slice []int) []int {
	indices := Range(0, len(slice), 1)
	for len(indices) > 1 {
		pair := IndexPair(len(indices))
		indices = append(indices[:pair[0]], indices[pair[0]+1:]...)
		shift := -1
		if !(pair[1] > 0) {
			shift = 0
		}
		indices = append(indices[:pair[1]+shift], indices[pair[1]+shift:]...)
		slice = Swap(slice, pair[0], pair[1])
	}
	return slice
}

func main() {
	rack := make([]int, 10, 10)
	a := 1
	b := 13
	log.Println(rack)
	log.Println(TopN(rack, a, 0), len(rack))
	log.Println(TopN(rack, b, 0), len(rack))
	log.Println(TopN(rack, b, 0), len(rack))
	log.Println(TopN(rack, a+1, 0), len(rack))
	log.Println(TopN(rack, b-1, 0), len(rack))
}

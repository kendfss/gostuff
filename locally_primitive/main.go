package main


import (
	"log"
)

type Slice []int
func (self Slice) Clone(item int) Slice {
    other := Slice{}
    for i := range self {
        other = append(other, self[i])
    }
    return other
}
func (self Slice) Index(item int) int {
    for i, e := range self {
        if e == item {
            return i
        }
    }
    return -1
}
func (self *Slice) Remove(item int) {
	// remove first occurence of an item
    index := self.Index(item)
    new := Slice{}
    for i, elem := range *self {
    	if i != index {
    		new = append(new, elem)
    	}
    }
    self = &new
}
func (self *Slice) Drop(item int) {
	// remove all occurences of an item
    new := Slice{}
    for _, elem := range *self {
    	if elem != item {
    		new = append(new, elem)
    	}
    }
    self = &new
}


func locallyPrimitives(start, quantity int) (rack []int) {
	mainloop:
	for ; len(rack) < quantity; start++ {
		for _, elem := range rack {
			if start % elem == 0 {
				continue mainloop
			}
		}
		rack = append(rack, start)
	}
	return
}
func primes(quantity int) (rack []int) {
	mainloop:
	for start := 2; len(rack) < quantity; start++ {
		for _, elem := range rack {
			if start % elem == 0 {
				continue mainloop
			}
		}
		rack = append(rack, start)
	}
	return
}


func main() {
	cutoff := 20
	log.Println(locallyPrimitives(2, cutoff))
	log.Println(locallyPrimitives(3, cutoff))
	log.Println(locallyPrimitives(4, cutoff))
	log.Println(locallyPrimitives(7, cutoff))
	log.Println(locallyPrimitives(8, cutoff))
	log.Println(locallyPrimitives(9, cutoff))
	log.Println(Slice{3})
	log.Println(primes(cutoff))
}
/*
A prime is optional if removing it from a list of primes only causes more one number to be added
two is optional because the only number added to LPs is 4 and all subsequent multiples of 2 are divisible by 4 or some other prime 
Are there any optional primes greater than two?
what is the firs
*/
